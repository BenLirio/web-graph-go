package frontier

import (
	"errors"
	"net/http"
	"fmt"
	"../links"
)

var MAX_DEPTH int = 3

func search(url string, depth int) error {
	if depth > MAX_DEPTH {
		return errors.New(fmt.Sprintf("Skipping %s – Depth: %d reached", url, depth))
	}
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New(fmt.Sprintf("Error Getting %s", url))
	}
	cLinks := make(chan string)
	go links.GetLinks(resp.Body, cLinks)
	for link := range cLinks {
		go search(link, depth + 1)
	}
	return nil
}
