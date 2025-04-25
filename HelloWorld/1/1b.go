package main

import "fmt"

func main() {
	x := 10
	y := 20
	var temp int
	temp = x
	x = y
	y = temp
	fmt.Println(x, y)
}
