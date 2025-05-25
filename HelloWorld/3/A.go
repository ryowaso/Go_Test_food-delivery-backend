package main

import "fmt"

func test(arr *[3]string) {
	arr[0] = "cuit"
}

func main() {
	var array [3]string = [3]string{"china", "sichuan", "chengdu"}
	fmt.Println("The array is:", array)
	test(&array)
	fmt.Println("The array is:", array)
}

//写出输出结果
