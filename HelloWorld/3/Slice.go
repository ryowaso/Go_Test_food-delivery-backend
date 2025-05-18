package main

import "fmt"

func main() {
	//var arr [5]int = [5]int{1, 2, 3, 4, 5}
	//slice := arr[:3]
	//fmt.Println(slice)
	//slice[1] = 100
	//fmt.Println(slice[1])
	//fmt.Println(slice[2])

	//var slice []int = []int{1, 2, 3}
	//fmt.Println("1. The value of slice is:", slice)
	//slice1 := append(slice, 4, 5)
	//fmt.Println("2. The value of slice1 is:", slice1)
	//slice = append(slice, 4, 5, 6)
	//fmt.Println("3. The value of slice is:", slice)
	//slice = append(slice, slice...)
	//fmt.Println("4. The value of slice is:", slice)

	//var slice []int = []int{1, 2, 3}
	//fmt.Println("1. The value of slice is:", slice) //
	//slice1 := append(slice, 4, 5)
	//fmt.Println("2. The value of slice1 is:", slice1) //
	//fmt.Printf("The address of slice is %p\n", &slice[0])
	//fmt.Printf("The address of slice1 is %p", &slice1[0])

	//var slice1 []int = []int{1, 2, 3, 4, 5}
	//var slice2 []int = make([]int, 10)
	//fmt.Println("The value of slice2", slice2) //
	//copy(slice1, slice2)
	//fmt.Println("The value of slice1 is ", slice1) //
	//fmt.Println("The value of slice2 is ", slice2) //

	var slice1 []int = []int{1, 2, 3, 4, 5}
	var slice2 []int = []int{3, 4, 5, 6}
	fmt.Println("The value of slice1", slice1)
	copy(slice1, slice2)
	fmt.Println("The value of slice1 is ", slice1)
	fmt.Println("The value of slice2 is ", slice2)

}
