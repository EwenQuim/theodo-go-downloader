package main

import (
	"fmt"
	"io"
	"os"
)

// progressWriter is a wrapper around a file and a source reader that
// implements the [io.Writer] interface. It's used to write the response
// body to a file and to send progress updates to the UI.
type progressWriter struct {
	url        string
	total      int
	downloaded int
	file       *os.File
	source     io.Reader
}

func (pw *progressWriter) Start() error {
	// TeeReader calls pw.Write() each time a new response is received
	_, err := io.Copy(pw.file, io.TeeReader(pw.source, pw))
	return err
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	pw.downloaded += len(p)

	if pw.total == 0 {
		return 0, fmt.Errorf("total size is unknown")
	}

	percentage := 100 * pw.downloaded / pw.total

	fmt.Println(pw.url, ": downloaded", percentage, "%")

	return len(p), nil
}
