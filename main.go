package main

import (
	"./links"
	"os"
	"fmt"
	"net/http"
	"sync"
)

type Edge struct {
	from string
	to   string
}

var visited sync.Map
func main() {
	i := 0
	if len(os.Args) < 2 {
		fmt.Println("Must specify a file to save to")
	}
	fp, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	queue := make(chan Edge)
	go search("https://google.com", queue)
	for edge := range queue {
		i += 1
		fmt.Println(i)
		fp.Write([]byte(fmt.Sprintf("%s\t%s\n", edge.from, edge.to)))
	}
}

func search(url string, queue chan Edge) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	cLinks := make(chan string)
	go links.GetLinks(resp.Body, cLinks)
	for link := range cLinks {
		if _, ok := visited.Load(link); ok == false {
			visited.Store(link, true)
			var edge Edge
			edge.from = url
			edge.to = link
			queue <- edge
		}
	}
}
