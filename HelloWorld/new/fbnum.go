package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	r := fboo(a - 1)
	fmt.Println(r)
}
func fboo(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fboo(n-1) + fboo(n-2)
}
