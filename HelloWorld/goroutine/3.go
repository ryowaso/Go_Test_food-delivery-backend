package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("a", i)
	}
}
func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("b", i)
	}
}
func c() {
	for i := 1; i < 10; i++ {
		fmt.Println("c", i)
	}
}
func main() {
	runtime.GOMAXPROCS(3)
	go a()
	go b()
	go c()
	time.Sleep(time.Second)
}
