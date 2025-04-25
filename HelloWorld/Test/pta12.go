package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	rr := test(s)
	fmt.Println(string(rr))
}

func test(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
