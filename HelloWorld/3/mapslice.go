package main

import "fmt"

func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Println(index, value)
	}
	fmt.Println("after init")
	//对切片中的初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "xxx"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "xx"
	for index, value := range mapSlice {
		fmt.Println(index, value)
	}
}
