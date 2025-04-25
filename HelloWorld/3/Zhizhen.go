package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("type of b:%T \n", b)
	c := *b
	a = 20
	fmt.Printf("type of c: %T \n", c)
	fmt.Printf("value of c %v \n", c)
	fmt.Printf("value of b %v \n", b)

}
