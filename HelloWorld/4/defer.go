package main

import "fmt"

func main() {
	x := 1
	y := 2
	defer calc("aa", x, calc("a", x, y))
	x = 10
	defer calc("bb", x, calc("b", x, y))
	y = 20

}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
