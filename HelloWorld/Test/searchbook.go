package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	ISBN  string
	Name  string
	Price float64
	Next  *Book
}

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Scan(&n)

		var head *Book
		var tail *Book

		maxPrice := -1.0

		for j := 0; j < n; j++ {
			var isbn string
			var name string
			var pricestr string
			fmt.Scan(&isbn, &name, &pricestr)

			price, _ := strconv.ParseFloat(pricestr, 64)

			book := &Book{ISBN: isbn, Name: name, Price: price}
			if head == nil {
				head = book
				tail = book
			} else {
				tail.Next = book
				tail = book
			}

			if price > maxPrice {
				maxPrice = price
			}
		}

		c := head
		for c != nil {
			if c.Price == maxPrice {
				fmt.Printf("%s %s %.2f\n", c.ISBN, c.Name, c.Price)
			}
			c = c.Next
		}
	}
}
