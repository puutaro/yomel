package sh

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	logGuard    = "####"
	labelPrefix = "#"
	redStart    = "\x1b[31m"
	redEnd      = "\x1b[0m"
)

func Exec(yomelInfos []YomelInfo) {
	numCmds := len(yomelInfos)
	if numCmds == 0 {
		return
	}

	cmds := make([]*exec.Cmd, numCmds)
	stdoutBuffers := make([]*bytes.Buffer, numCmds)
	stderrBuffers := make([]*bytes.Buffer, numCmds)

	var nextStdin io.Reader = os.Stdin

	// 1. Build the pipeline structure
	for i, yomelInfo := range yomelInfos {
		cmd := exec.Command("bash", "-c", yomelInfo.CmdStrs)
		cmds[i] = cmd

		cmd.Stdin = nextStdin
		stdoutBuffers[i] = new(bytes.Buffer)
		stderrBuffers[i] = new(bytes.Buffer)
		cmd.Stderr = stderrBuffers[i]

		stdoutPipe, _ := cmd.StdoutPipe()
		// Forward data to the next command while simultaneously writing to its own log buffer
		nextStdin = io.TeeReader(stdoutPipe, stdoutBuffers[i])
	}

	// Immediately start consuming data from the final downstream in the background (asynchronously)
	lastCmdDone := make(chan struct{})
	// 2. Start consuming data in the background
	go func() {
		_, _ = io.Copy(os.Stdout, nextStdin)
		close(lastCmdDone)
	}()
	// 3. Start all commands simultaneously
	for _, cmd := range cmds {
		_ = cmd.Start()
	}

	// 4. Wait for all data to flow through (consumption to finish)
	<-lastCmdDone

	// 5. Wait for each command process itself to terminate
	cmdHasError := false
	for _, cmd := range cmds {
		if err := cmd.Wait(); err == nil {
			continue
		}
		cmdHasError = true
	}

	// 6. Finally, output decorated logs to os.Stderr based on flag conditions
	for i, yomelInfo := range yomelInfos {
		stdoutLen := stdoutBuffers[i].Len()

		shouldLogStdout := yomelInfo.IsLog && stdoutLen > 0

		if !cmdHasError && !shouldLogStdout {
			continue
		}
		printDecoratedLog(
			yomelInfo.No,
			yomelInfo.Desc,
			yomelInfo.CmdStrs,
			yomelInfo.LogFilter,
			yomelInfo.ErrLogFilter,
			stderrBuffers[i],
			stdoutBuffers[i],
			cmdHasError,
		)
	}
}

func printDecoratedLog(
	no int,
	desc,
	cmdName string,
	logFilterShell string,
	errLogFilterShell string,
	stderrBuf,
	stdoutBuf *bytes.Buffer,
	cmdHasError bool,
) {

	title := fmt.Sprintf("%s [%d]YOMEL LOG", logGuard, no)

	fmt.Fprintf(
		os.Stderr,
		"%s\n%s stage: \n%s\n%s cmd: \n%s\n",
		title,
		labelPrefix,
		desc,
		labelPrefix,
		cmdName,
	)

	if cmdHasError {
		write2Std(
			os.Stderr,
			fmt.Sprintf(
				"%s%s error:%s\n",
				labelPrefix,
				redStart,
				redEnd,
			),
			stderrBuf,
			errLogFilterShell,
		)
	}
	write2Std(
		os.Stderr,
		fmt.Sprintf("%s stdout:\n", labelPrefix),
		stdoutBuf,
		logFilterShell,
	)

	fmt.Fprintf(os.Stderr, "%s\n\n", logGuard)
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
	fmt.Fprintf(w, "%sfilter shell err:%s\n", redStart, redEnd)
	w.Write(filterShellCmdStderrBuf.Bytes())
	addNewline(w, filterShellCmdStderrBuf)
}

func addNewline(w io.Writer, buf *bytes.Buffer) {
	if buf.Bytes()[buf.Len()-1] == '\n' {
		return
	}
	fmt.Fprintln(w)
}
