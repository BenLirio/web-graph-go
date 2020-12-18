package main

import (
	"os"
	"fmt"
	"golang.org/x/net/html"
	"./search"
)

var visited map[string]bool
var fp *os.File

func main() {
	var err error
	visited = make(map[string]bool)
	fp, err = os.Create("data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go walkWeb("http://google.com")
	for {}
}

func walkWeb(url string) {
	fp.Write([]byte(fmt.Sprintf("%s\n", url)))
	//fmt.Println(url)
	htmlReader, err := search.GetHTMLReader(url)
	if err != nil {
		//fmt.Printf("ERROR Getting %s\n", url)
		return
	}
	rootNode, err := html.Parse(htmlReader)
	if err != nil {
		//fmt.Printf("ERROR Parsing %s\n", url)
		return
	}
	links, err := search.GetLinks(rootNode)
	if err != nil {
		//fmt.Printf("ERROR Finding links in %s\n", url)
		return
	}
	for _, link := range links {
		go walkWeb(link)
//		if _, ok := visited[link]; ok != true {
//			visited[link] = false
//			go walkWeb(link)
//		}
	}
}
