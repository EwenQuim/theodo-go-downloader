package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one URL")
		os.Exit(1)
	}
	requests(os.Args[1:])
}

func requests(urls []string) error {
	for _, url := range urls {
		fmt.Println(url)
	}
	return nil
}
