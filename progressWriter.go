package main

import (
	"io"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var program *tea.Program

// progressWriter is a wrapper around a file and a source reader that
// implements the [io.Writer] interface. It's used to write the response
// body to a file and to send progress updates to the UI.
type progressWriter struct {
	total      int
	downloaded int
	file       *os.File
	source     io.Reader
}

func (pw *progressWriter) Start() {
	// TeeReader calls pw.Write() each time a new response is received
	_, err := io.Copy(pw.file, io.TeeReader(pw.source, pw))
	if err != nil {
		program.Send(progressErrMsg{err})
	}
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	pw.downloaded += len(p)
	if pw.total > 0 {
		program.Send(progressMsg(float64(pw.downloaded) / float64(pw.total)))
	}
	return len(p), nil
}
