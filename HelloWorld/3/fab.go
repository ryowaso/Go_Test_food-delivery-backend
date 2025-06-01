package main

import "fmt"

func Fab(n int) int {

	//if n == 1 || n == 2 {
	//	return 1
	//} else {
	//	return Fab(n-1) + Fab(n-2)
	//}

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b

}
func generateslice(n int) []int {

	var slice []int
	for i := 1; i <= n; i++ {
		slice = append(slice, Fab(i))
	}
	fmt.Println(slice)
	return slice

}
func main() {
	slice := generateslice(10)
	for index, value := range slice {
		fmt.Println(index, value)
	}
}
