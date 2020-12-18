package search

import (
	"golang.org/x/net/html"
	"net/http"
	"io"
)

func GetHTMLReader(url string) (r io.ReadCloser, err error){
	resp, err := http.Get(url)
	if err != nil {
		return r, err
	}
	return resp.Body, nil

}

func GetLinks(root *html.Node) ([]string, error) {
	var links []string
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if len(attr.Val) >= 4 && attr.Val[:4] == "http" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(root)
	return links, nil
}
