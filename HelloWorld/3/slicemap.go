package main

import "fmt"

func main() {
	var slicemap = make(map[string][]string, 3)
	fmt.Println(slicemap)
	fmt.Println("after init")
	key := "a"
	value, ok := slicemap[key]
	if ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "b", "c")
	slicemap[key] = value
	fmt.Println(slicemap)
}
