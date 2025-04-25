package main

import "fmt"

var cache = make(map[string]string)

func main() {
	fmt.Println(cacheGet("a1"))
	fmt.Println(cacheGet("a2"))
	fmt.Println(cacheGet("a3"))
}
func cacheGet(key string) string {
	v, ok := cache[key]
	if ok {
		return v
	}
	data := cacheGet("a1" + key)
	//data:="data1"+ key
	cache[key] = string(data)
	return data
}
