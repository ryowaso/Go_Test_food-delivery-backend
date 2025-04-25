package main

import "fmt"

func main() {
	var n int = 10
	fmt.Printf("前%d个fb数是%d\n", n, fb(n))

}
func fb(n int) []int {
	if n <= 0 {
		return []int{}
	}
	fbo := make([]int, n)
	fbo[0] = 0
	if n > 1 {
		fbo[1] = 1
	}
	for i := 2; i < n; i++ {
		fbo[i] = fbo[i-1] + fbo[i-2]

	}
	return fbo
}
