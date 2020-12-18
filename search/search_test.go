package search

import (
	"golang.org/x/net/html"
	"testing"
)

func TestGetHTML(t *testing.T) {
	htmlReader, err := GetHTMLReader("http://google.com")
	defer htmlReader.Close()
	if err != nil {
		t.Error(err)
	}
	node, err := html.Parse(htmlReader)
	if err != nil {
		t.Error(err)
	}
	if node.Type != html.DocumentNode {
		t.Error("Root node should be a documnet node")
	}
}

func TestGetLinks(t *testing.T) {
	htmlReader, err := GetHTMLReader("http://google.com")
	defer htmlReader.Close()
	if err != nil {
		t.Error(err)
	}
	node, err := html.Parse(htmlReader)
	if err != nil {
		t.Error(err)
	}
	links, err := GetLinks(node)
	if err != nil {
		t.Error(err)
	}
	for _, link := range links {
		println(link)
	}
}
