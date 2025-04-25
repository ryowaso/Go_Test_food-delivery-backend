package main

import "fmt"

func main() {
	x := 2
	foo(&x)
	fmt.Println(x)
}

func foo(y *int) {
	*y = *y * 2
}
