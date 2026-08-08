package sh

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
)

const (
	redStart       = "\x1b[31m"
	colorEnd       = "\x1b[39m"
	blueGreenStart = "\x1b[30m"
	ansiEnd        = "\x1b[0m"
	boldStart      = "\x1b[1m"
	boldEnd        = "\x1b[22m"
	underlineStart = "\x1b[4m"
	underlineEnd   = "\x1b[24m"
)

type yomelLog struct {
	yomelInfo      YomelInfo
	stageInfos     []StageInfo
	stdoutBuffers  []*bytes.Buffer
	stderrBuffers  []*bytes.Buffer
	cmdHasError    bool
	startTime      time.Time
	stageEndTimes  []time.Time
	stageDurations []time.Duration
}

func Exec(yomelInfo YomelInfo) int {
	stageInfos := yomelInfo.StageInfos
	numCmds := len(stageInfos)
	if numCmds == 0 {
		return ExitSuccess
	}
	if yomelInfo.IsGen {
		totalPipeCmdStr := makeTotalPipeCmd(
			stageInfos,
			func(stInfo StageInfo) string {
				return stInfo.CmdStrs
			},
		)
		outputCmd(totalPipeCmdStr)
		return ExitSuccess
	}
	if yomelInfo.IsDirect {
		totalPipeCmdStr := makeTotalPipeCmd(
			stageInfos,
			func(stInfo StageInfo) string {
				return stInfo.CmdStrs
			},
		)
		return directExec(totalPipeCmdStr)
	}

	cmds := make([]*exec.Cmd, numCmds)
	stdoutBuffers := make([]*bytes.Buffer, numCmds)
	stderrBuffers := make([]*bytes.Buffer, numCmds)

	var nextStdin io.Reader = os.Stdin

	var wg sync.WaitGroup
	// 1. Build the pipeline structure
	for i, stageInfo := range stageInfos {
		cmd := exec.Command("bash", "-c", stageInfo.CmdStrs)
		cmds[i] = cmd

		cmd.Stdin = nextStdin
		stdoutBuffers[i] = new(bytes.Buffer)
		stderrBuffers[i] = new(bytes.Buffer)
		stderrPipe, _ := cmd.StderrPipe()
		stderrTee := io.TeeReader(stderrPipe, stderrBuffers[i])
		wg.Add(1)
		go func() {
			defer wg.Done()
			if yomelInfo.IsLiveStdErr {
				_, _ = io.Copy(os.Stderr, stderrTee)
			} else {
				_, _ = io.Copy(io.Discard, stderrTee)
			}
		}()

		stdoutPipe, _ := cmd.StdoutPipe()
		// Forward data to the next command while simultaneously writing to its own log buffer
		nextStdin = io.TeeReader(stdoutPipe, stdoutBuffers[i])
	}

	// Immediately start consuming data from the final downstream in the background (asynchronously)
	lastCmdDone := make(chan struct{})
	// 2. Start consuming data in the background
	go func() {
		if yomelInfo.IsLiveStdout {
			_, _ = io.Copy(os.Stdout, nextStdin)
		} else {
			_, _ = io.Copy(io.Discard, nextStdin)
		}
		close(lastCmdDone)
	}()
	// 3. Start all commands simultaneously
	for _, cmd := range cmds {
		_ = cmd.Start()
	}

	// 4. Wait for all data to flow through (consumption to finish)
	<-lastCmdDone
	wg.Wait()

	// 5. Wait for each command process itself to terminate
	startTime := time.Now()
	stageDurations := make([]time.Duration, numCmds)
	stageEndTimes := make([]time.Time, numCmds)
	cmdHasError := false
	var exitCode int = 0
	for i, cmd := range cmds {
		cmdErr := cmd.Wait()
		stageEndTimes[i] = time.Now()
		stageDurations[i] = stageEndTimes[i].Sub(startTime)
		if cmdErr == nil {
			continue
		}
		cmdHasError = true
		exitCode = extractErrCode(cmdErr)
	}

	// 6. Finally, output decorated logs to os.Stderr based on flag conditions
	yl := yomelLog{
		yomelInfo:      yomelInfo,
		stageInfos:     stageInfos,
		stdoutBuffers:  stdoutBuffers,
		stderrBuffers:  stderrBuffers,
		cmdHasError:    cmdHasError,
		startTime:      startTime,
		stageEndTimes:  stageEndTimes,
		stageDurations: stageDurations,
	}
	combinedLog := yl.make()

	if combinedLog.Len() > 0 {
		_, _ = os.Stderr.Write(combinedLog.Bytes())
	}
	if cmdHasError {
		return exitCode
	}
	return ExitSuccess
}

func (yl *yomelLog) make() bytes.Buffer {
	yomelTitle := yl.yomelInfo.Title
	totalPipeCmdStr := makeTotalPipeCmd(
		yl.stageInfos,
		func(stInfo StageInfo) string {
			return stInfo.CmdStrsWithComment
		},
	)
	var combinedLog bytes.Buffer
	stageLen := len(yl.stageInfos)

	// find first log output index
	firstLogIdx := -1
	for i, stageInfo := range yl.stageInfos {
		shouldLog := stageInfo.IsLog || yl.cmdHasError
		if yl.cmdHasError || shouldLog {
			firstLogIdx = i
			break
		}
	}
	if firstLogIdx == -1 {
		return combinedLog
	}

	yl.printYomelLogStartHolder(&combinedLog)
	if stageLen > 1 {
		yl.printTitleLog(&combinedLog, yomelTitle)
		yl.printTotalCmd(&combinedLog, totalPipeCmdStr)
	} else {
		fmt.Fprintf(
			&combinedLog,
			"\n",
		)
	}
	for i, stageInfo := range yl.stageInfos {
		shouldLog := stageInfo.IsLog || yl.cmdHasError
		if !yl.cmdHasError && !shouldLog {
			continue
		}
		yl.printDecoratedLog(
			&combinedLog,
			stageInfo,
			i,
			yl.stderrBuffers[i],
			yl.stdoutBuffers[i],
			shouldLog,
		)
	}
	return combinedLog
}

func extractErrCode(err error) int {
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		return exitError.ExitCode()
	}
	return ExitErrGeneral
}

func outputCmd(totalPipeCmdStr string) {
	fmt.Fprintln(
		os.Stdout,
		totalPipeCmdStr,
	)
}

func directExec(totalPipeCmdStr string) int {
	if len(totalPipeCmdStr) == 0 {
		return ExitSuccess
	}
	cmd := exec.Command("bash", "-c", totalPipeCmdStr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return extractErrCode(err)
	}
	return ExitSuccess
}

func makeTotalPipeCmd(
	stageInfos []StageInfo,
	cmdStrFn func(stageInfo StageInfo) string,
) string {
	var cmdPipeline string
	for i, info := range stageInfos {
		cmdStr := cmdStrFn(info)
		if i == 0 {
			cmdPipeline = cmdStr
		} else {
			cmdPipeline = fmt.Sprintf("%s \\\n| %s", cmdPipeline, cmdStr)
		}
	}
	return cmdPipeline
}

func (yl *yomelLog) printYomelLogStartHolder(
	w io.Writer,
) {
	var yomelLogStartHolderBuffer bytes.Buffer
	yomelLogStartHolderBuffer.WriteString("\n")
	yomelLogStartHolderBuffer.WriteString(
		fmt.Sprintf(
			"%s%s%s_%s%s%s\n",
			underlineStart,
			boldStart,
			"Yomel-log",
			convertTimeStampStr(yl.startTime),
			boldEnd,
			underlineEnd,
		),
	)
	w.Write(
		yomelLogStartHolderBuffer.Bytes(),
	)
}

func (yl *yomelLog) printTitleLog(
	w io.Writer,
	title string,
) {
	if title == "" {
		return
	}
	fmt.Fprintf(
		w,
		"%s\n%s\n\n",
		convertUnderAndFirstBold("Title"),
		capitalizeFirst(title),
	)
}

func (yl *yomelLog) printTotalCmd(
	w io.Writer,
	totalPipeCmdStr string,
) {
	fmt.Fprintf(
		w,
		"%s\n%s\n\n",
		convertUnderAndFirstBold("Total-cmd"),
		formatSideComment(totalPipeCmdStr),
	)
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

func convertUnderAndFirstBold(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return underlineStart + boldStart + string(r) + boldEnd + s[size:] + underlineEnd
}

func (yl *yomelLog) printDecoratedLog(
	w io.Writer,
	stageInfo StageInfo,
	index int,
	stderrBuf,
	stdoutBuf *bytes.Buffer,
	shouldLog bool,
) {
	duration := yl.stageDurations[index]
	durationStr := fmt.Sprintf("+%.6fs", duration.Seconds())
	endTime := yl.stageEndTimes[index]
	stageHeader := fmt.Sprintf(
		"%s[%d]_%s(%s)",
		"Stage",
		stageInfo.No,
		convertTimeStampStr(endTime),
		durationStr,
	)
	fmt.Fprintf(
		w,
		"%s\n%s\n\n%s \n%s\n\n",
		convertUnderAndFirstBold(
			stageHeader,
		),
		capitalizeFirst(stageInfo.Desc),
		convertUnderAndFirstBold("Cmd"),
		formatSideComment(stageInfo.CmdStrsWithComment),
	)
	if shouldLog {
		yl.write2Std(
			w,
			makeNormalOrRedStdErrLabel(yl.cmdHasError),
			stderrBuf,
			stageInfo.ErrLogFilter,
		)
	}
	yl.write2Std(
		w,
		fmt.Sprintf("%s\n", convertUnderAndFirstBold("Stdout")),
		stdoutBuf,
		stageInfo.LogFilter,
	)
}

func makeNormalOrRedStdErrLabel(hasErr bool) string {
	logGenre := "Progress"
	if hasErr {
		logGenre = "Error"
		return fmt.Sprintf(
			"%s%s%s\n",
			redStart,
			convertUnderAndFirstBold(logGenre),
			colorEnd,
		)
	}
	return fmt.Sprintf(
		"%s\n",
		convertUnderAndFirstBold(logGenre),
	)
}

func (yl *yomelLog) write2Std(w io.Writer, label string, buf *bytes.Buffer, filterShell string) {
	if buf.Len() <= 0 {
		return
	}
	fmt.Fprint(w, label)
	if filterShell == "" {
		w.Write(buf.Bytes())
		addNewline(w, buf)
		return
	}
	filterShellCmd := exec.Command("bash", "-c", filterShell)
	filterShellCmd.Stdin = buf
	filterShellCmdStdoutBuf := new(bytes.Buffer)
	filterShellCmd.Stdout = filterShellCmdStdoutBuf
	filterShellCmdStderrBuf := new(bytes.Buffer)
	filterShellCmd.Stderr = filterShellCmdStderrBuf
	if err := filterShellCmd.Run(); err == nil || filterShellCmdStderrBuf.Len() <= 0 {
		w.Write(filterShellCmdStdoutBuf.Bytes())
		addNewline(w, filterShellCmdStdoutBuf)
		return
	}
	fmt.Fprintf(
		w,
		"%s\n",
		fmt.Sprintf(
			"%s%s%s",
			redStart,
			convertUnderAndFirstBold("filter shell err"),
			colorEnd,
		),
	)
	w.Write(filterShellCmdStderrBuf.Bytes())
	addNewline(w, filterShellCmdStderrBuf)
}

func addNewline(w io.Writer, buf *bytes.Buffer) {
	if buf.Len() == 0 {
		return
	}
	length := buf.Len()
	bufBytes := buf.Bytes()
	if bufBytes[length-1] == '\n' {
		return
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w)
}

func formatSideComment(cmdStr string) string {
	sideCommentBlank := domain.SideCommentBlank

	lines := strings.Split(cmdStr, "\n")
	var maxBaseRune int
	baseRunes := make([]int, len(lines))
	idxs := make([]int, len(lines))
	for i, line := range lines {
		idx := strings.Index(line, sideCommentBlank)
		idxs[i] = idx
		if idx == -1 {
			continue
		}
		baseLine := line[:idx]
		rNum := utf8.RuneCountInString(baseLine)
		baseRunes[i] = rNum
		if rNum > maxBaseRune {
			maxBaseRune = rNum
		}
	}
	var sb strings.Builder
	sb.Grow(len(cmdStr) + len(lines)*4)
	concatBlank := "    "
	for i, line := range lines {
		idx := idxs[i]
		if idx == -1 {
			sb.WriteString(line)
			if i < len(lines)-1 {
				sb.WriteByte('\n')
			}
			continue
		}

		baseLine := line[:idx]
		commentLine := line[idx:]
		diff := maxBaseRune - baseRunes[i]

		sb.WriteString(baseLine)
		if diff > 0 {
			for d := 0; d < diff; d++ {
				sb.WriteByte(' ')
			}
		}
		sb.WriteString(concatBlank)

		commentBody := commentLine[len(sideCommentBlank):]
		sb.WriteString(commentBody)

		if i < len(lines)-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
func convertTimeStampStr(t time.Time) string {
	return t.Format("2006/01/02-15:04:05.000000")
}
