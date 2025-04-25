package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	f := fb(n)
	for _, v := range f {
		fmt.Printf("%d ", v)
	}
}
func fb(n int) []int {
	if n <= 0 {
		return []int{}
	}
	fbo := make([]int, n)
	fbo[0] = 1
	if n > 1 {
		fbo[1] = 1
	}
	for i := 2; i < n; i++ {
		fbo[i] = fbo[i-1] + fbo[i-2]

	}
	return fbo
}
