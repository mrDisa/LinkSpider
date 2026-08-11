package main

import (
    "flag"
    "fmt"
    "net/url"
    "linkspider/internal/checker"
    "linkspider/internal/fetcher"
    "linkspider/internal/parser"
)

func main() {
	urlPtr := flag.String("url", "", "URL to check")
	flag.Parse()
	pageURL := *urlPtr

	if pageURL == "" {
		fmt.Println("Ошибка: нужно указать -url")
		return
	}

	reader, err := fetcher.FetchHTML(pageURL)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	links, err := parser.ExtractLinks(reader)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	base, err := url.Parse(pageURL)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	for _, rawLink := range links {
		ref, err := url.Parse(rawLink)
		if err != nil {
			continue
		}
		resolved := base.ResolveReference(ref)
		resolvedURL := resolved.String()
		link := checker.CheckLink(resolvedURL)
		
		if link.Error != "" {
			fmt.Printf("Link - %s - STATUS: %d [FAIL]: %s\n", link.URL, link.StatusCode, link.Error)
		} else if link.Alive {
			fmt.Printf("Link - %s - STATUS: %d [OK]\n", link.URL, link.StatusCode)
		} else {
			fmt.Printf("Link - %s - STATUS: %d [BROKEN]\n", link.URL, link.StatusCode)
		}
	}
}