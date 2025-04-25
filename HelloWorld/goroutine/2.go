package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hello1(i int) {
	defer wg.Done()
	fmt.Println("hello world", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Hello1(i)
	}
	wg.Wait()
}
