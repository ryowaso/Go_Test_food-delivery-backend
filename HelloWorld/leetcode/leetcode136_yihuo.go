package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
func main() {
	// 示例测试
	nums := []int{2, 2, 1}
	fmt.Println("只出现一次的数字是:", singleNumber(nums)) // 输出: 1

	nums = []int{4, 1, 2, 1, 2}
	fmt.Println("只出现一次的数字是:", singleNumber(nums)) // 输出: 4

	nums = []int{1}
	fmt.Println("只出现一次的数字是:", singleNumber(nums)) // 输出: 1
}
