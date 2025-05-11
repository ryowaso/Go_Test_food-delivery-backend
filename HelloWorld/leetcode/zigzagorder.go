package main

import "fmt"

type TreeNodeZ struct {
	Val   int
	Left  *TreeNodeZ
	Right *TreeNodeZ
}

func main() {
	root := &TreeNodeZ{
		Val: 1,
		Left: &TreeNodeZ{
			Val:   2,
			Left:  &TreeNodeZ{Val: 4},
			Right: &TreeNodeZ{Val: 5},
		},
		Right: &TreeNodeZ{
			Val:   3,
			Right: &TreeNodeZ{Val: 6},
		},
	}

	result := zigzagOrder(root)
	fmt.Println(result) // 输出：[[1] [3 2] [4 5 6]]
}
func zigzagOrder(root *TreeNodeZ) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []*TreeNodeZ{root}
	leftToRight := true //方向控制，起始从左到右

	for len(queue) > 0 {
		n := len(queue)
		vals := make([]int, n)

		for i := range queue {
			node := queue[0]
			queue = queue[1:]
			//根据当前层的方向决定插入位置
			if leftToRight {
				vals[i] = node.Val
			} else {
				vals[n-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, vals)
		leftToRight = !leftToRight //切换方向
	}
	return res
}
