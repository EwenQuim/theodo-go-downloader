package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Write the response body to the file
	_, err = file.Write(body)
	if err != nil {
		return err
	}

	fmt.Println(url, ": wrote to file", file.Name())
	return nil
}
