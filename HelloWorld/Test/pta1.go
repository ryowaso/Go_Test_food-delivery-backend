package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i * i
	}
	fmt.Println(sum)

}

//func isEven(n int) bool {
//	return n%2 == 0
//}
