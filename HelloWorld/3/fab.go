package main

import "fmt"

func Fab(n int) int {

	if n == 1 || n == 2 {
		return 1
	} else {
		return Fab(n-1) + Fab(n-2)
	}
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
