package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var x int
	fmt.Scan(&x)

	i, count := binarySearch(a, x)
	fmt.Println(i)
	fmt.Println(count)
}
func binarySearch(a []int, x int) (int, int) {
	l, r := 0, len(a)-1
	count := 0
	for l <= r {
		mid := l + (r-l)/2
		count++
		if a[mid] == x {
			return mid, count
		} else if a[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1, count
}
