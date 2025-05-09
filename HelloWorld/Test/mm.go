package main

import "fmt"

func main() {
	var c, n int
	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		if is(n) {
			fmt.Println(po(n))
		} else {
			fmt.Println(0)
		}
	}
}

func is(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func po(n int) int {
	if n < 2 {
		return 0
	}
	count := 1
	for i := 3; i <= n; i += 2 {
		if is(i) {
			count++
		}
	}
	return count
}
