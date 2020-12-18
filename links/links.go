package links

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"regexp"
	"errors"
)

func parseLink(val string) (string, error) {
	r := regexp.MustCompile("^https?")
	protocol := r.Find([]byte(val))
	if len(protocol) == 0 {
		return "", errors.New("Value does not have a valid protocol")
	}
	r = regexp.MustCompile("//([^/]*)")
	domain := r.FindSubmatch([]byte(val))
	if len(domain) < 2 {
		return "", errors.New("Value does not havea  valid domain")
	}
	return string(protocol) + "://" + string(domain[1]), nil
}

func GetLinks(htmlReader io.Reader, c chan string) {
	tokenizer := html.NewTokenizer(htmlReader)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if token.DataAtom != atom.A {
			continue
		}
		if token.Type != html.StartTagToken {
			continue
		}
		for _, attr := range token.Attr {
			if attr.Key != "href" {
				continue
			}
			if link, err := parseLink(attr.Val); err == nil {
				c <- link
			}
		}
	}
	close(c)
}
