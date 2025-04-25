package main

import "fmt"

func main() {
	a := 3
	b := 4
	x1 := &a
	x2 := &b
	fmt.Println(a, b)
	swap(x1, x2)
	fmt.Println(a, b)
}
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
