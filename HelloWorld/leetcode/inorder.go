package main

import "fmt"

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func main() {
	root := &TreeNode1{
		Val: 1,
		Left: &TreeNode1{
			Val:   2,
			Left:  &TreeNode1{Val: 4},
			Right: &TreeNode1{Val: 5},
		},
		Right: &TreeNode1{Val: 3},
	}
	result := inorder(root)
	fmt.Println(result)
}

// 迭代，加上指针
func inorder(root *TreeNode1) []int {
	stack := []*TreeNode1{}
	var vals []int
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		vals = append(vals, cur.Val)
		cur = cur.Right
	}
	return vals
}
