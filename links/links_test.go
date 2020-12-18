package links

import (
	"testing"
	"os"
	"net/http"
	"strings"
)

var samplePage string = `
<a href="http://google.com">to google</a>
`
func TestLinksOnStaticPage(t *testing.T) {
	pageReader := strings.NewReader(samplePage)
	c := make(chan string)
	go GetLinks(pageReader, c)
	for range c {}
}

func BenchmarkGetLinks(b *testing.B) {
	fp, err := os.Open("testdata/index.html")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := make(chan string)
		fp.Seek(0,0)
		go GetLinks(fp, c)
		for range c {}
	}
}

func TestLinksOnLivePage(t *testing.T) {
	t.Skip()
	resp, err := http.Get("http://google.com")
	if err != nil {
		t.Error(err)
	}
	c := make(chan string)
	go GetLinks(resp.Body, c)
	for range c {}
}
