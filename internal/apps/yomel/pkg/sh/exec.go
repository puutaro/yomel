package sh

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	logGuard       = "####"
	labelPrefix    = "#"
	redStart       = "\x1b[31m"
	blueGreenStart = "\x1b[30m"
	colorEnd       = "\x1b[0m"
)

func Exec(yomelInfo YomelInfo) error {
	stageInfos := yomelInfo.StageInfos
	numCmds := len(stageInfos)
	if numCmds == 0 {
		return nil
	}
	if yomelInfo.IsGen {
		outputCmd(stageInfos)
		return nil
	}
	if yomelInfo.IsDirect {
		directExec(stageInfos)
		return nil
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
	cmdHasError := false
	for _, cmd := range cmds {
		if err := cmd.Wait(); err == nil {
			continue
		}
		cmdHasError = true
	}

	// 6. Finally, output decorated logs to os.Stderr based on flag conditions
	yomelTitle := yomelInfo.Title
	for i, stageInfo := range stageInfos {
		// stdoutLen := stdoutBuffers[i].Len()

		shouldLog := stageInfo.IsLog || cmdHasError

		if !cmdHasError && !shouldLog {
			continue
		}
		if i == 0 {
			printTitleLog(yomelTitle)
		}
		printDecoratedLog(
			stageInfo.No,
			stageInfo.Desc,
			stageInfo.CmdStrs,
			stageInfo.LogFilter,
			stageInfo.ErrLogFilter,
			stderrBuffers[i],
			stdoutBuffers[i],
			shouldLog,
			cmdHasError,
		)
	}
	if cmdHasError {
		return errors.New("failed")
	}
	return nil
}
func outputCmd(stageInfo []StageInfo) {
	fmt.Fprintln(
		os.Stdout,
		makeDirectCmd(stageInfo),
	)
}
func directExec(stageInfos []StageInfo) {
	if len(stageInfos) == 0 {
		return
	}
	cmdPipeline := makeDirectCmd(stageInfos)
	cmd := exec.Command("bash", "-c", cmdPipeline)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%serror: pipeline failed: %v%s\n", redStart, err, colorEnd)
	}
}
func makeDirectCmd(stageInfos []StageInfo) string {
	var cmdPipeline string
	for i, info := range stageInfos {
		if i == 0 {
			cmdPipeline = info.CmdStrs
		} else {
			cmdPipeline = fmt.Sprintf("%s \\\n| %s", cmdPipeline, info.CmdStrs)
		}
	}
	return cmdPipeline
}

func printTitleLog(
	title string,
) {
	if title == "" {
		return
	}
	boldStart := "\x1b[1m"
	titleHolder := fmt.Sprintf("%s%s YOMEL-TITLE:%s", boldStart, logGuard, colorEnd)
	titleSentence := boldStart + capitalizeFirst(title) + colorEnd
	fmt.Fprintf(
		os.Stderr,
		"%s\n%s\n\n",
		titleHolder,
		titleSentence,
	)
}
func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
func printDecoratedLog(
	no int,
	desc,
	cmdName string,
	logFilterShell string,
	errLogFilterShell string,
	stderrBuf,
	stdoutBuf *bytes.Buffer,
	shouldLog bool,
	cmdHasError bool,
) {

	timestamp := time.Now().Format("2006/01/02-15:04:05.000000")
	title := fmt.Sprintf("%s YOMEL-LOG[%d]_%s:", logGuard, no, timestamp)

	fmt.Fprintf(
		os.Stderr,
		"%s\n%s Stage: \n%s\n\n%s Cmd: \n%s\n\n",
		title,
		labelPrefix,
		capitalizeFirst(desc),
		labelPrefix,
		cmdName,
	)

	if shouldLog {
		write2Std(
			os.Stderr,
			makeNormalOrRedStdErrLabel(cmdHasError),
			stderrBuf,
			errLogFilterShell,
		)
	}
	write2Std(
		os.Stderr,
		fmt.Sprintf("%s Stdout:\n", labelPrefix),
		stdoutBuf,
		logFilterShell,
	)

	if len(stdoutBuf.Bytes()) > 0 {
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func makeNormalOrRedStdErrLabel(hasErr bool) string {
	logGenre := "Progress"
	if hasErr {
		logGenre = "Error"
		return fmt.Sprintf(
			"%s%s %s:%s\n",
			labelPrefix,
			redStart,
			logGenre,
			colorEnd,
		)
	}
	return fmt.Sprintf(
		"%s %s:\n",
		labelPrefix,
		logGenre,
	)
}

func write2Std(w io.Writer, label string, buf *bytes.Buffer, filterShell string) {
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
	fmt.Fprintf(w, "%sfilter shell err:%s\n", redStart, colorEnd)
	w.Write(filterShellCmdStderrBuf.Bytes())
	addNewline(w, filterShellCmdStderrBuf)
}

func addNewline(w io.Writer, buf *bytes.Buffer) {
	if buf.Len() == 0 {
		return
	}
	if buf.Bytes()[buf.Len()-1] == '\n' {
		return
	}
	fmt.Fprintln(w)
}
