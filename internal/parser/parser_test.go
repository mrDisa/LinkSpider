package parser

import (
	"strings"
	"testing"
	"reflect"
)

func TestExtractLinks(t * testing.T) {
	htmlInput := `<html><body><a href="/foo">x</a><a href="https://bar.com">y</a></body></html>`
	case1, err := ExtractLinks(strings.NewReader(htmlInput)) 
	expected := []string{"/foo", "https://bar.com"}

	if err != nil {
		t.Fatalf("error - %s", err)
	}

	if !reflect.DeepEqual(case1, expected) {
		t.ErrorF("ExtractLinks(html with href) = %v; expected %v", case1, expected)
	}
}