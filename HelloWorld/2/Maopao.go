package main

import "fmt"

func main() {
	arr := []int{5, 2, 1, 3, 4}
	n := len(arr)
	var temp int

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	fmt.Println("排序后的数组：")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
