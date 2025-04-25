package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello ", name)
	var a int = 1
	var b, c int = 2, 3
	var (
		d int    = 4
		e int    = 5
		f string = "hello"
	)
	g := "world"

	fmt.Println(a, b, c, d, e, f, g)

	fmt.Println(add(a, b))

	fmt.Println(moreresult(a, b))

	fmt.Println(namefunc(a, b))

	fmt.Println(increment)

	//var x int
	//fmt.Scanln(&x)
	//switch {
	//case x > 1:
	//	fmt.Println("x is greater than 1")
	//case x < -1:
	//	fmt.Println("x is less than -1")
	//default:
	//	fmt.Println("x is equal to -1")
	//}

	i := 0
	for i < 10 {
		i++
		if i == 5 {
			continue
		}
		fmt.Printf("%d ", i, "hi")
	}

	var p [5]int = [5]int{1, 2, 3, 4, 5}
	o := [...]int{1, 2, 3, 4, 5}
	fmt.Println(p)
	fmt.Println(p[0], p[1], p[2])
	fmt.Println(o[0], o[1], o[2])
}

func add(a int, b int) int {
	return a + b
}

func moreresult(a int, b int) (int, int) {
	return a + b, a - b
}

func namefunc(a, b int) (p int) {
	p = a * b
	return
}

func increment(i int) {
	i++
}
