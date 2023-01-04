package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one URL")
		os.Exit(1)
	}
	err := requests(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func requests(urls []string) error {
	for _, url := range urls {
		err := fetch(url)
		if err != nil {
			return fmt.Errorf("error fetching %s: %v", url, err)
		}
	}
	return nil
}

func fetch(url string) error {
	fmt.Println(url, ": fetching...")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	fmt.Println(url, ":", resp.Status)

	filename := filepath.Base(url)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	fmt.Println(url, ": created file", file.Name())

	progress := progressWriter{
		url:    url,
		total:  int(resp.ContentLength),
		file:   file,
		source: resp.Body,
	}

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(url, ": your download has started in the background and will be opened when complete")
	}()

	err = progress.Start()
	if err != nil {
		return err
	}

	fmt.Println(url, ": wrote to file", file.Name())

	open.Run(file.Name())
	return nil
}
