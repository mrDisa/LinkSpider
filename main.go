package main

import (
    "fmt"
    "io"
    "linkspider/internal/fetcher"
)

func main() {
	reader, err := fetcher.FetchHTML("https://example.com")
	if err != nil {
		fmt.Println("error: ", err)
	}
	data, _ := io.ReadAll(reader)
	fmt.Println(string(data))
}