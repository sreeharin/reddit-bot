package main

import (
	"bufio"
	"os"
    "github.com/sreeharin/reddit-bot/crawler"
)

func main() {
	file_name := "sub_reddits.txt"
	if len(os.Args) > 1 {
		file_name = os.Args[1]
	}

	f, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
        crawler.Crawl(scanner.Text())
	}
}
