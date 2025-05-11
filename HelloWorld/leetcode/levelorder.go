package main

import "fmt"

type TreeNodeQ struct {
	Val   int
	Left  *TreeNodeQ
	Right *TreeNodeQ
}

// 队列
func levelOrder(root *TreeNodeQ) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	queue := []*TreeNodeQ{root}
	for len(queue) > 0 {
		n := len(queue)
		vals := make([]int, n)
		for i := range vals {
			node := queue[0]
			queue = queue[1:]
			vals[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
func main() {
	root := &TreeNodeQ{
		Val: 1,
		Left: &TreeNodeQ{
			Val:   2,
			Left:  &TreeNodeQ{Val: 4},
			Right: &TreeNodeQ{Val: 5},
		},
		Right: &TreeNodeQ{
			Val:   3,
			Right: &TreeNodeQ{Val: 6},
		},
	}
	result := levelOrder(root)
	fmt.Println(result) // 输出 [[1] [2 3] [4 5 6]]
}
