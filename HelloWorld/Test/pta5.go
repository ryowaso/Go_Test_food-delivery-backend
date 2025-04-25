package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])
	}
	sum := 0
	count := 0
	for _, v := range slice {
		sum += v
		if v > 90 {
			count++
		}
	}
	fmt.Printf("sum=%d\n", sum)
	fmt.Printf("count=%d\n", count)
}
