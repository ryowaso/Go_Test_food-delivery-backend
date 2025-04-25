package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanln(&a, &b)
	sum := a + b
	fmt.Println(sum)
}
