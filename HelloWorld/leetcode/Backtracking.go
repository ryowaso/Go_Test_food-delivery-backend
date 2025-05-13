package main

import (
	"fmt"
	"sort"
)

//以leetcode2094为例

func findEvenNumbers(digits []int) []int {
	n := len(digits)
	used := make([]bool, n)
	resultSet := make(map[int]bool)

	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == 3 {
			// 跳过前导零
			if path[0] == 0 {
				return
			}
			// 最后一位必须是偶数
			if path[2]%2 != 0 {
				return
			}
			// 拼接成数字
			num := path[0]*100 + path[1]*10 + path[2]
			resultSet[num] = true
			return
		}

		for i := 0; i < n; i++ {
			// 每个元素只能用一次
			if used[i] {
				continue
			}
			// 为了避免重复，跳过同一层中重复的数字
			if i > 0 && digits[i] == digits[i-1] && !used[i-1] {
				continue
			}

			used[i] = true
			path = append(path, digits[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	// 排序以便剪枝去重
	sort.Ints(digits)
	backtrack([]int{})

	// 提取并排序结果
	result := make([]int, 0, len(resultSet))
	for num := range resultSet {
		result = append(result, num)
	}
	sort.Ints(result)
	return result
}
func main() {
	digits := []int{1, 2, 3}
	result := findEvenNumbers(digits)
	fmt.Println(result) // 示例输出: [132 312]
}
