package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)

		books := make([][3]string, n)
		maxPrice := -1.0

		for j := 0; j < n; j++ {
			var isbn, name, pricep string
			fmt.Scan(&isbn, &name, &pricep)

			price, _ := strconv.ParseFloat(pricep, 64)
			books[j] = [3]string{isbn, name, pricep}

			if price > maxPrice {
				maxPrice = price
			}

		}
		for _, book := range books {
			price, _ := strconv.ParseFloat(book[2], 64)

			if price == maxPrice {
				fmt.Printf("%s %s %s\n", book[0], book[1], book[2])
			}
		}
	}

}
