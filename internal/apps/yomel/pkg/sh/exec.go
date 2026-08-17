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

	"github.com/charmbracelet/x/term"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/domain"
	"github.com/puutaro/yomel/internal/apps/yomel/pkg/env"
)

const (
	redStart       = "\x1b[31m"
	colorEnd       = domain.ColorEnd
	blueGreenStart = "\x1b[30m"
	boldStart      = "\x1b[1m"
	boldEnd        = "\x1b[22m"
	underlineStart = "\x1b[4m"
	underlineEnd   = "\x1b[24m"
)

const (
	sectionPrefixAndLastNewlineNum        = 2
	miniSectionOrContentsPrefixNewlineNum = 1
	sectionSentenceSeparator              = " "
)

type yomelLog struct {
	yomelInfo        YomelInfo
	stageInfos       []StageInfo
	stdoutBuffers    []*bytes.Buffer
	stderrBuffers    []*bytes.Buffer
	cmdHasError      bool
	startTime        time.Time
	stageDurations   []time.Duration
	isTerminal       bool
	isLightColorMode bool
}

func Exec(
	yomelInfo YomelInfo,
	yomelEnv env.YomelEnv,
) int {
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
		yomelInfo:        yomelInfo,
		stageInfos:       stageInfos,
		stdoutBuffers:    stdoutBuffers,
		stderrBuffers:    stderrBuffers,
		cmdHasError:      cmdHasError,
		startTime:        startTime,
		stageDurations:   stageDurations,
		isTerminal:       yomelEnv.IsTerminal,
		isLightColorMode: yomelEnv.IsLightColorMode,
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
	isLightColorMode := yl.isLightColorMode
	titleStartBackgroundColor := outputTitleStartColor(
		yl.yomelInfo.BackgroundColor,
		yl.isTerminal,
		isLightColorMode,
	)
	yl.printYomelLogStartHolder(
		&combinedLog,
		titleStartBackgroundColor,
	)
	if stageLen > 1 {
		yl.printTitleLog(
			&combinedLog,
			yomelTitle,
			titleStartBackgroundColor,
		)
		yl.printTotalCmd(
			&combinedLog,
			totalPipeCmdStr,
			titleStartBackgroundColor,
		)
	}
	var curSectionColor string
	for i, stageInfo := range yl.stageInfos {
		shouldLog := stageInfo.IsLog || yl.cmdHasError
		if !yl.cmdHasError && !shouldLog {
			continue
		}
		backgroundColor := stageInfo.BackgroundColor
		curSectionColor = outputSectionColorStart(
			backgroundColor,
			yl.isTerminal,
			isLightColorMode,
			curSectionColor,
		)
		yl.printDecoratedLog(
			&combinedLog,
			stageInfo,
			i,
			yl.stderrBuffers[i],
			yl.stdoutBuffers[i],
			shouldLog,
			curSectionColor,
		)
	}
	fmt.Fprint(&combinedLog, compNewLine(&combinedLog, sectionPrefixAndLastNewlineNum))
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

func colorizePipelineCmd(
	line string,
	fgColor string,
	isTerminal bool,
) string {
	// すでに空行の場合や、カラーコードが不要な場合はそのまま返す
	if line == "" || fgColor == "" || !isTerminal {
		return line
	}
	lineWithColorEndTrim := strings.ReplaceAll(line, colorEnd, "")
	lines := strings.Split(lineWithColorEndTrim, "\n")
	for i, line := range lines {
		if !strings.HasSuffix(line, "\\") {
			lines[i] = fgColor + line + colorEnd
			continue
		}
		base := line[:len(line)-1]
		lines[i] = fgColor + base + fgColor + "\\" + colorEnd
	}
	return strings.Join(lines, "\n")
}

func (yl *yomelLog) printYomelLogStartHolder(
	w io.Writer,
	titleStartBackgroundColor string,
) {
	var yomelLogStartHolderBuffer bytes.Buffer
	yomelLogStartHolderBuffer.WriteString("\n")
	logStartRaw := fmt.Sprintf(
		"%s%s%s",
		"Yomel-log",
		sectionSentenceSeparator,
		convertTimeStampStr(yl.startTime),
	)
	isTerminal := yl.isTerminal
	logStartWithColor := makeForegroundColor(
		logStartRaw,
		yl.yomelInfo.ForegroundColor,
		isTerminal,
	)
	logStartHolderWithDeco := ""
	switch isTerminal {
	case true:
		logStartHolderWithDeco = makeFullWidthBgColorLine(
			fmt.Sprintf(
				"%s%s%s%s%s",
				underlineStart,
				boldStart,
				logStartWithColor,
				boldEnd,
				underlineEnd,
			),
			titleStartBackgroundColor,
			isTerminal,
		)
	default:
		logStartHolderWithDeco = logStartWithColor
	}
	yomelLogStartHolderBuffer.WriteString(
		logStartHolderWithDeco,
	)
	w.Write(
		yomelLogStartHolderBuffer.Bytes(),
	)
}

func (yl *yomelLog) printTitleLog(
	w io.Writer,
	title string,
	titleStartColor string,
) {
	if title == "" {
		return
	}
	newLIneStr := compNewLine(w, miniSectionOrContentsPrefixNewlineNum)
	foregroundColor := yl.yomelInfo.ForegroundColor
	isTerminal := yl.isTerminal
	titleHOlderWithFgColorUnder := makeForegroundColor(
		convertUnderAndFirstBold(
			"Title",
			isTerminal,
		),
		foregroundColor,
		isTerminal,
	)
	titleHolderWithDeco := makeFullWidthBgColorLine(
		titleHOlderWithFgColorUnder,
		titleStartColor,
		isTerminal,
	)
	titleCon := makeForegroundColor(
		capitalizeFirst(title),
		foregroundColor,
		isTerminal,
	)
	titleHolderAndCon := fmt.Sprintf(
		"%s\n%s",
		titleHolderWithDeco,
		titleCon,
	)
	fmt.Fprintf(
		w,
		"%s%s",
		newLIneStr,
		titleHolderAndCon,
	)
}

func (yl *yomelLog) printTotalCmd(
	w io.Writer,
	totalPipeCmdStr string,
	titleStartColor string,
) {
	newlineStr := compNewLine(w, sectionPrefixAndLastNewlineNum)
	isTerminal := yl.isTerminal
	yomelInfo := yl.yomelInfo
	foregroundColor := yomelInfo.ForegroundColor
	totalCmdSectionWithFgColorUnsder := makeForegroundColor(
		convertUnderAndFirstBold(
			"Total-cmd",
			isTerminal,
		),
		foregroundColor,
		isTerminal,
	)
	totalCmdSectionWithDeco := makeFullWidthBgColorLine(
		totalCmdSectionWithFgColorUnsder,
		titleStartColor,
		isTerminal,
	)
	titleCommentColorStart := yomelInfo.TitleCommentForegroundColor
	fmt.Fprintf(
		w,
		"%s%s\n%s",
		newlineStr,
		totalCmdSectionWithDeco,
		colorizePipelineCmd(
			formatSideComment(
				decideCommentColor(
					totalPipeCmdStr,
					titleCommentColorStart,
					isTerminal,
				),
			),
			foregroundColor,
			isTerminal,
		),
	)
}

func decideCommentColor(cmdWithCome, commentRepColor string, isTerminal bool) string {
	if !isTerminal ||
		cmdWithCome == "" ||
		commentRepColor == "" {
		return cmdWithCome
	}
	return strings.ReplaceAll(cmdWithCome, domain.GrayStart, commentRepColor)
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}

func convertUnderAndFirstBold(s string, isTerminal bool) string {
	if s == "" {
		return ""
	}
	if !isTerminal {
		return s
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
	curSectionColor string,
) {
	foreroundColor := stageInfo.ForegroundColor
	duration := yl.stageDurations[index]
	durationStr := fmt.Sprintf("+%.6fs", duration.Seconds())
	newlineCompedForStage := compNewLine(w, sectionPrefixAndLastNewlineNum)
	stageHeader := fmt.Sprintf(
		"%s[%d]%s%s",
		"Stage",
		stageInfo.No,
		sectionSentenceSeparator,
		durationStr,
	)
	isTerminal := yl.isTerminal
	stageHeaderWithUnerFirstBold := makeForegroundColor(
		convertUnderAndFirstBold(
			stageHeader,
			isTerminal,
		),
		foreroundColor,
		isTerminal,
	)
	stageHeaderWithDeco := makeFullWidthBgColorLine(
		stageHeaderWithUnerFirstBold,
		curSectionColor,
		isTerminal,
	)
	descWithFgColorCapitalizeFirst := makeForegroundColor(
		capitalizeFirst(stageInfo.Desc),
		foreroundColor,
		isTerminal,
	)
	stageCon := fmt.Sprintf(
		"%s\n%s",
		stageHeaderWithDeco,
		descWithFgColorCapitalizeFirst,
	)
	fmt.Fprintf(
		w,
		"%s%s",
		newlineCompedForStage,
		stageCon,
	)
	newlineCompedForCmd := compNewLine(w, sectionPrefixAndLastNewlineNum)
	cmdSectionWithFgColorUnder := makeForegroundColor(
		convertUnderAndFirstBold(
			"Cmd",
			isTerminal,
		),
		foreroundColor,
		isTerminal,
	)
	cmdSectionWithDeco := makeBackgoundColor(
		cmdSectionWithFgColorUnder,
		curSectionColor,
		isTerminal,
	)
	curCommentColorStart := outputCommentColorStart(
		stageInfo.CommentColorStart,
		yl.yomelInfo.CommentColor,
		isTerminal,
	)
	cmdBodyWithDeco := colorizePipelineCmd(
		formatSideComment(
			decideCommentColor(
				stageInfo.CmdStrsWithComment,
				curCommentColorStart,
				isTerminal,
			),
		),
		foreroundColor,
		isTerminal,
	)
	fmt.Fprintf(
		w,
		"%s%s \n%s",
		newlineCompedForCmd,
		cmdSectionWithDeco,
		cmdBodyWithDeco,
	)
	stdoutHolderColor := foreroundColor
	if yl.cmdHasError {
		stdoutHolderColor = redStart
	}
	stdErrSectionWithFgColorUnserBold := makeForegroundColor(
		makeNormalOrRedStdErrLabel(
			yl.cmdHasError,
			isTerminal,
		),
		stdoutHolderColor,
		isTerminal,
	)
	stdErrSectionWithDeco := makeBackgoundColor(
		stdErrSectionWithFgColorUnserBold,
		curSectionColor,
		isTerminal,
	)
	if shouldLog {
		newlineForStdErr := compNewLine(w, miniSectionOrContentsPrefixNewlineNum)
		yl.write2Std(
			w,
			newlineForStdErr+
				stdErrSectionWithDeco,
			stderrBuf,
			stageInfo.ErrLogFilter,
			foreroundColor,
			isTerminal,
		)
	}
	newlineCompedForStdout := compNewLine(w, sectionPrefixAndLastNewlineNum)
	stdoutHolderWithFgColorBoldUnder := makeForegroundColor(
		convertUnderAndFirstBold(
			"Stdout",
			isTerminal,
		),
		foreroundColor,
		isTerminal,
	)
	stdoutHolderWithDeco := makeBackgoundColor(
		stdoutHolderWithFgColorBoldUnder,
		curSectionColor,
		isTerminal,
	)
	yl.write2Std(
		w,
		fmt.Sprintf(
			"%s%s",
			newlineCompedForStdout,
			stdoutHolderWithDeco,
		),
		stdoutBuf,
		stageInfo.LogFilter,
		foreroundColor,
		isTerminal,
	)
}

func makeNormalOrRedStdErrLabel(hasErr bool, isTerminal bool) string {
	logGenre := "Progress"
	if hasErr {
		logGenre = "Error"
	}
	return fmt.Sprintf(
		"%s",
		convertUnderAndFirstBold(
			logGenre,
			isTerminal,
		),
	)
}

func (yl *yomelLog) write2Std(
	w io.Writer,
	label string,
	buf *bytes.Buffer,
	filterShell string,
	foregroundColor string,
	isTerminal bool,
) {
	if buf.Len() <= 0 {
		return
	}
	fmt.Fprint(w, label)
	if filterShell == "" {
		newLineStr := compNewLine(w, miniSectionOrContentsPrefixNewlineNum)
		w.Write([]byte(newLineStr))
		w.Write(
			bufStrToByteWithFgColor(
				buf.String(),
				foregroundColor,
				isTerminal,
			),
		)
		return
	}
	filterShellCmd := exec.Command("bash", "-c", filterShell)
	filterShellCmd.Stdin = buf
	filterShellCmdStdoutBuf := new(bytes.Buffer)
	filterShellCmd.Stdout = filterShellCmdStdoutBuf
	filterShellCmdStderrBuf := new(bytes.Buffer)
	filterShellCmd.Stderr = filterShellCmdStderrBuf
	if err := filterShellCmd.Run(); err == nil || filterShellCmdStderrBuf.Len() <= 0 {
		newLineStr := compNewLine(w, miniSectionOrContentsPrefixNewlineNum)
		w.Write([]byte(newLineStr))
		w.Write(
			bufStrToByteWithFgColor(
				filterShellCmdStdoutBuf.String(),
				foregroundColor,
				isTerminal,
			),
		)
		return
	}
	newlineForFilterShell := compNewLine(w, miniSectionOrContentsPrefixNewlineNum)
	filterShellErrStrWithDeco := fmt.Sprintf(
		"%s%s%s",
		redStart,
		convertUnderAndFirstBold(
			"filter shell err",
			yl.isTerminal,
		),
		colorEnd,
	)
	fmt.Fprintf(
		w,
		"%s%s",
		newlineForFilterShell,
		filterShellErrStrWithDeco,
	)
	newLineStrForFiltershellErrCon := compNewLine(w, miniSectionOrContentsPrefixNewlineNum)
	w.Write([]byte(newLineStrForFiltershellErrCon))
	w.Write(
		bufStrToByteWithFgColor(
			filterShellCmdStderrBuf.String(),
			foregroundColor,
			isTerminal,
		),
	)
}

func bufStrToByteWithFgColor(
	bufStr string,
	foregroundColor string,
	isTerminal bool,
) []byte {
	bufStrWithFgColor := makeForegroundColor(
		bufStr,
		foregroundColor,
		isTerminal,
	)
	return []byte(bufStrWithFgColor)
}

func formatSideComment(cmdStr string) string {
	sideCommentBlank := domain.SideCommentBlank
	lines := strings.Split(cmdStr, "\n")
	var maxBaseRune int
	baseRunes := make([]int, len(lines))
	idxs := make([]int, len(lines))
	for i, lineSrc := range lines {
		idx := strings.Index(lineSrc, sideCommentBlank)
		idxs[i] = idx
		if idx == -1 {
			continue
		}
		baseLine := lineSrc[:idx]
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

func compNewLine(w io.Writer, newLIneNum int) string {
	newlineStr := "\n"
	buf, ok := w.(*bytes.Buffer)
	if !ok {
		return strings.Repeat(newlineStr, newLIneNum)
	}
	b := buf.Bytes()
	newlineCount := 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != '\n' {
			break
		}
		newlineCount++
	}
	addCount := newLIneNum - newlineCount
	if addCount >= 0 {
		return strings.Repeat("\n", addCount)
	}
	return ""
}

func outputTitleStartColor(
	backGroundColor string,
	isTerm bool,
	isLightColorMode bool,
) string {
	if !isTerm {
		return ""
	}
	if backGroundColor != "" {
		return backGroundColor
	}
	ligthTarcoizBakcgroundStart := "\x1b[48;5;193m"
	if isLightColorMode {
		return ligthTarcoizBakcgroundStart
	}
	deepbrownBackgroundStart := "\x1b[48;5;52m"
	return deepbrownBackgroundStart
}
func outputSectionColorStart(
	backgroudColorStartStr string,
	isTerm bool,
	isLightColorMode bool,
	beforeColor string,
) string {
	if !isTerm {
		return ""
	}
	if backgroudColorStartStr != "" {
		return backgroudColorStartStr
	}
	switch isLightColorMode {
	case true:
		lightBlueBakcgroundStart := "\x1b[48;5;159m"
		// lightGreenBackgroundStart := "\x1b[48;5;121m"
		lightGreenBackgroundStart := "\x1b[48;5;157m"
		if beforeColor == lightBlueBakcgroundStart {
			return lightGreenBackgroundStart
		} else {
			return lightBlueBakcgroundStart
		}
	default:
		darkGreenBackgroundStart := "\x1b[48;5;22m"
		darkBlueBakcgroundStart := "\x1b[48;5;18m"
		if beforeColor == darkGreenBackgroundStart {
			return darkBlueBakcgroundStart
		} else {
			return darkGreenBackgroundStart
		}
	}
}
func outputCommentColorStart(
	stageCommentColorStart string,
	ctrlCommentColorStart string,
	isTerm bool,
) string {
	if !isTerm {
		return ""
	}
	if stageCommentColorStart != "" {
		return stageCommentColorStart
	}
	return ctrlCommentColorStart
}

func makeFullWidthBgColorLine(
	text string,
	colorStart string,
	isTerminal bool,
) string {
	if !isTerminal {
		return text
	}
	widthPtr := getTerminalWidth(os.Stderr)
	if widthPtr == nil {
		return makeBackgoundColor(
			text,
			colorStart,
			isTerminal,
		)
	}
	backgroundColorEnd := "\x1b[49m"
	plainText := stripAnsi(text)
	visLen := utf8.RuneCountInString(plainText)
	// ターミナルの幅に満たない分の空白をパディング
	paddingLen := *widthPtr - visLen
	if paddingLen < 0 {
		paddingLen = 0
	}
	// 元のテキスト（装飾付き）の末尾に、足りない分のスペースを追加
	paddedText := text + fmt.Sprintf("%*s", paddingLen, "")
	if colorStart == "" {
		return paddedText
	}
	return fmt.Sprintf("%s%s%s", colorStart, paddedText, backgroundColorEnd)
}
func makeBackgoundColor(
	str string,
	colorStart string,
	isTerminal bool,
) string {
	if colorStart == "" || !isTerminal {
		return str
	}
	backgroundColorEnd := "\x1b[49m"
	return fmt.Sprintf(
		"%s%s%s",
		colorStart,
		str,
		backgroundColorEnd,
	)
}
func makeForegroundColor(
	str string,
	colorStart string,
	isTerminal bool,
) string {
	if colorStart == "" || !isTerminal {
		return str
	}
	return fmt.Sprintf(
		"%s%s%s",
		colorStart,
		str,
		colorEnd,
	)
}

func getTerminalWidth(f *os.File) *int {
	// f.Fd() を使ってターミナルのファイル記述子を渡す
	width, _, err := term.GetSize(f.Fd())
	if err != nil {
		return nil // 取得失敗時は一般的な80文字幅をデフォルトにする
	}
	return &width
}

func stripAnsi(s string) string {
	var buf strings.Builder
	inEscape := false
	for i := 0; i < len(s); {
		if i+1 < len(s) && s[i] == 0x1b && s[i+1] == '[' {
			inEscape = true
			i += 2
			continue
		}
		if inEscape {
			c := s[i]
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
				inEscape = false
			}
			i++
			continue
		}
		buf.WriteByte(s[i])
		i++
	}
	return buf.String()
}
