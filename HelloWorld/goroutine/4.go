package main

import (
	"fmt"
	"runtime"
	"sync"
)

func a1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("a", i)
	}
}
func b1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("b", i)
	}
}
func c1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("c", i)
	}
}
func main() {
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	wg.Add(3)
	go a1(&wg)
	go b1(&wg)
	go c1(&wg)
	wg.Wait()
}
