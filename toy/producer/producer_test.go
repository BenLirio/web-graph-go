package producer

import (
	"testing"
	"fmt"
)

func TestProduction(t *testing.T) {
	c := make(chan string)
	p := NewProducer(c)
	go p.generate()
	for product := range c {
		fmt.Println(product)
	}
}
