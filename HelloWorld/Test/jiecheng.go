package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(factorial(a))
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
