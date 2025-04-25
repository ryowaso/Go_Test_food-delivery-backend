package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a+b > c && a+c > b && b+c > a {
		p := (a + b + c) / 2
		area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
		fmt.Printf("Output:\n%.2f\n", area)
	} else {
		fmt.Printf("Output:\ndata error\n")
	}
}
