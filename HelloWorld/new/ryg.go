package main

import (
	"fmt"
	"time"
)

func main() {

	str := "str"
label:
	for {
		switch str {
		case "red":
			fmt.Println("red")
			time.Sleep(2 * time.Second)
			str = "red"
		case "yellow":
			fmt.Println("yellow")
			time.Sleep(3 * time.Second)
			str = "yellow"
		case "green":
			fmt.Println("green")
			time.Sleep(5 * time.Second)
			str = "green"
		}
		goto label
	}
}
