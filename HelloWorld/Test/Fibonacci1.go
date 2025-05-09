package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	result := []int{0, 1}
	for i := 2; i < n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}
	return result
}

func main() {
	n := 5
	fmt.Println(fibonacci(n))
}
