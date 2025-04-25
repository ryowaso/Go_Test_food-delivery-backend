package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0
	f := 1
	for i := 0; i <= n; i++ {
		if i != 0 {
			f *= i
		}
		sum += f
	}
	fmt.Println(sum)
}
