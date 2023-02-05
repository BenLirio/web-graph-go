package frontier


import (
	"testing"
)

func TestSearch(t *testing.T) {
	go search("https://google.com", 0)
	for {}
}
