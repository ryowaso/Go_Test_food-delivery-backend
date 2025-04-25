package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var p1, p2 string
		fmt.Scan(&p1, &p2)
		if p1 == p2 {
			fmt.Println("TIE")
			continue
		}
		switch p1 {
		case "R":
			if p2 == "S" {
				fmt.Println("Player1")
			} else {
				fmt.Println("Player2")
			}
		case "S":
			if p2 == "P" {
				fmt.Println("Player1")
			} else {
				fmt.Println("Player2")
			}
		case "P":
			if p2 == "R" {
				fmt.Println("Player1")
			} else {
				fmt.Println("Player2")
			}
		default:
			fmt.Println("TIE")
		}
	}
}
