package main

import "fmt"

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func postorder(root *TreeNode2) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode2{root}
	var vals []int
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append([]int{node.Val}, vals...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return vals
}
func main() {
	root := &TreeNode2{
		Val: 1,
		Left: &TreeNode2{
			Val:   2,
			Left:  &TreeNode2{Val: 4},
			Right: &TreeNode2{Val: 5},
		},
		Right: &TreeNode2{Val: 3},
	}
	result := postorder(root)
	fmt.Println(result)
}
