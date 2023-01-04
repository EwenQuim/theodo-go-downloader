package main

import (
	"fmt"
	"net/http"
	"os"
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
	return nil
}
