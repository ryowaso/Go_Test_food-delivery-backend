package main

import "fmt"

// 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	bottom := triangle[len(triangle)-1]
	dp := make([]int, len(bottom))
	for i := range dp {
		dp[i] = bottom[i]
	}
	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}
