package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/skratchdot/open-golang/open"
)

func execute() error {
	url := flag.String("url", "", "url for the file to download")
	flag.Parse()

	if *url == "" {
		return errors.New("url is required")
	}

	resp, err := http.Get(*url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := filepath.Base(*url)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	pw := progressWriter{
		total:  int(resp.ContentLength),
		file:   file,
		source: resp.Body,
	}

	// Start Bubble Tea UI
	program = NewTUI(pw)

	// Start the download
	go pw.Start()

	// Run the UI
	_, err = program.Run()
	if err != nil {
		return err
	}

	open.Run(filename)

	return nil
}

func main() {
	err := execute()
	if err != nil {
		log.Fatal(err)
	}
}
