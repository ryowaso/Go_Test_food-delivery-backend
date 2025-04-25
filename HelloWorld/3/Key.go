package main

import "fmt"

func main() {
	userinfo := make(map[string]string, 8)
	userinfo["jack"] = "123456"
	userinfo["rose"] = "A123456"
	v, ok := userinfo["jack"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("无此用户")
	}
}
