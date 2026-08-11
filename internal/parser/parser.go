package parser

import (
    "golang.org/x/net/html"
    "io"
)

func ExtractLinks(htmlContent io.Reader) ([]string, error) {
    links := make([]string, 0)
	var crawler func(*html.Node)
    crawler = func(node *html.Node) {
        if node.Type == html.ElementNode && node.Data == "a" {
            for _, attr := range node.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
        }
        for child := node.FirstChild; child != nil; child = child.NextSibling {
            crawler(child)
        }
    }

	doc, err := html.Parse(htmlContent)
	if err != nil {
		return nil, err
	}

    crawler(doc)
	return links, nil
}