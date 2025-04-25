package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	sum := 0.0
	for i := 1; i <= N; i++ {
		S := math.Pow(-1, float64(i+1)) / (4*float64(i) - 3)
		sum += S
	}
	fmt.Printf("sum = %.6f\n", sum)
}
