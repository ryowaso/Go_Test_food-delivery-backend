package main

import "fmt"

func main() {
	var numRows int
	fmt.Scan(&numRows)
	fmt.Println(generate(numRows))
	var rowIndex int
	fmt.Scan(&rowIndex)
	fmt.Println(getRow(rowIndex))
}

func generate(numRows int) [][]int {
	c := make([][]int, numRows)
	for i := range c {
		c[i] = make([]int, i+1)
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
	return c
}

func getRow(rowIndex int) []int {
	t := generate(rowIndex + 1)
	return t[rowIndex]
}
