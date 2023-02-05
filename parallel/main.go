package main

import (
	"fmt"
)

var n int = 100

func leaders(follower chan bool, leader chan bool) {
	for i := 0; i < n; i++ {
		leader<-true
		<-follower
		fmt.Printf("%dL\n", i)
	}
}

func followers(follower chan bool, leader chan bool) {
	for i := 0; i < n; i++ {
		follower<-true
		<-leader
		fmt.Printf("%dF\n", i)
	}
}

func main() {
	follower := make(chan bool, 1)
	leader := make(chan bool, 1)
	go leaders(follower, leader)
	go followers(follower, leader)
	for {}
}
