package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	f := p(n)
	fmt.Println(f)
}

func p(n int) []int {
	f := make([]int, 0)

	if n < 2 {
		return f
	}

	for n%2 == 0 {
		f = append(f, 2)
		n /= 2
	}

	for i := 3; i < n; i += 2 {
		for n%i == 0 {
			f = append(f, i)
			n /= i
		}
	}

	if n > 2 {
		f = append(f, n)
	}

	return f
}
