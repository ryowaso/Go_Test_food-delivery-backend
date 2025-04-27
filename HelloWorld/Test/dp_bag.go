package main

import "fmt"

func main() {
	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}
	capacity := 4
	maxValue := knapsack2(weights, values, capacity)
	fmt.Println(maxValue)
}

func knapsack(weights []int, values []int, W int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if j < weights[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return dp[n][W]
}

//dp总结：每次面对一个物品，选择：
//
//不要它，价值继承上一行。
//
//要它，看看剩下空间能不能加上它的价值。
//
//取两种选择中更好的。

// 精简
func knapsack2(weights []int, values []int, W int) int {
	dp := make([]int, W+1)
	for i := 0; i < len(weights); i++ { //遍历每一个物品，i是物品编号
		for j := W; j >= weights[i]; j-- { //j是背包容量，从大到小，核心 避免自己更新自己
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
	}
	return dp[W]
}

// 滚动版
func knapsack3(weights []int, values []int, W int) int {
	n := len(weights)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= n; i++ {
		now := i % 2       //当前是哪一层
		pre := (i - 1) % 2 //上一层是哪一层
		for j := 0; j <= W; j++ {
			if j < weights[i-1] {
				dp[now][j] = dp[pre][j]
			} else {
				dp[now][j] = max(dp[pre][j], dp[pre][j-weights[i-1]]+values[i-1])
			}
		}
	}
	return dp[n%2][W]
}

// 多重背包,枚举
func knapsack4(weights []int, values []int, W int) int {
	n := len(weights)
	dp := make([]int, W+1)
	counts := make([]int, n+1)
	for i := 0; i < n; i++ { // 遍历每一个物品
		for j := W; j >= 0; j-- { // 背包容量从大到小
			for k := 1; k <= counts[i] && k*weights[i] <= j; k++ {
				// 枚举选0个、1个、2个...最多counts[i]个
				dp[j] = max(dp[j], dp[j-k*weights[i]]+k*values[i])
			}
		}
	}
	return dp[W]
}
