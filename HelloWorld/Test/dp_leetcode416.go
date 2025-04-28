package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true // 背包容量为0，什么都不拿也是ok的

	for _, num := range nums {
		for j := target; j >= num; j-- { // 倒序遍历！！保证每个数字只用一次
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}
func main() {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true},         // 可以分割为 [1,5,5] 和 [11]
		{[]int{1, 2, 3, 5}, false},         // 无法分割
		{[]int{1, 2, 3, 4, 5, 6, 7}, true}, // 可以分割为 [1,2,3,7] 和 [4,5,6]
		{[]int{2, 2, 3, 5}, false},         // 总和为12，但无法分割
		{[]int{1, 1}, true},                // 可以分割为 [1] 和 [1]
	}

	for i, tc := range testCases {
		result := canPartition(tc.nums)
		if result == tc.expected {
			fmt.Printf("测试用例 %d 通过: %v -> %v\n", i+1, tc.nums, result)
		} else {
			fmt.Printf("测试用例 %d 失败: %v -> 得到 %v, 期望 %v\n", i+1, tc.nums, result, tc.expected)
		}
	}
}
