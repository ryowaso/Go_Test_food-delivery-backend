package main

import (
	"fmt"
	"time"
)

func main() {
	go Hello()
	fmt.Println("start")
}

func Hello() {
	fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}
