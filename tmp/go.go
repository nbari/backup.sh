package main

import (
	"crypto"
	"time"
)

var c = make(chan int, 5)

func main() {
	go worker(1)
	for i := 0; i < 10; i++ {
		c <- i
		println(i)
	}
}

var xwww string

func worker(id int) {
	for {
		x := <-c
		print("---->", x)
		time.Sleep(time.Second)
	}
}
