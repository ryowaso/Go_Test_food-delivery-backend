package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		target := rand.Intn(100) + 1
		fmt.Println("a")
		for {
			var n int
			fmt.Println("b")
			fmt.Scan(&n)
			if n == target {
				fmt.Println("right")
				break
			} else if n > target {
				fmt.Println("big")
			} else if n < target {
				fmt.Println("small")
			}

		}
		var s string
		fmt.Print("yes or no")
		fmt.Scan(&s)
		if s == "yes" {
			fmt.Println("right")
			break
		}
	}
}
