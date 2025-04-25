package main

import "fmt"

const N = 1000000

func main() {
	var C, N int
	fmt.Scan(&C)
	for i := 0; i < C; i++ {
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}
		if isPrime(N) {
			fmt.Println(positionPrimes(N))
		} else {
			fmt.Println(0)
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func positionPrimes(p int) int {
	if p < 2 {
		return 0
	}
	count := 1
	for i := 3; i <= p; i += 2 {
		if isPrime(i) {
			count++
		}
	}
	return count
}

//从3开始，每次增加2（只检查奇数，因为大于2的偶数都不是素数）
