package main

import "fmt"

func main() {
	a := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, a1 := range a {
		for _, a2 := range a1 {
			//fmt.Printf("%s\t", a2)
			fmt.Println(a2)
		}
		fmt.Println()
	}
}
