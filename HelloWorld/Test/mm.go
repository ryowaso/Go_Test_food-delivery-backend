package main

import "fmt"

func main() {

	var intArr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := intArr[1:3]
	fmt.Println("The slice is: ", slice)
	fmt.Println("The len of slice is: ", len(slice))
	fmt.Println("The cap of slice is: ", cap(slice))
}
