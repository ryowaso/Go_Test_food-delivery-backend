package main

import "fmt"

func main() {
	length := 5
	width := 3
	area := length * width
	perimeter := (length + area) * 2
	fmt.Println(perimeter)
	fmt.Println(area)
}
