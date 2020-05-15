package main

import "math"

type TreeNode struct {
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

func main() {

}

//func isBalenced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	return math.Abs(float64(maxDepth(root.Left)) - float64(maxDepth(root.Right))) <= 1 && isBalenced(root.Left) && isBalenced(root.Right)
//}
//
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	} else {
//		left := maxDepth(root.Left)
//		right := maxDepth(root.Right)
//		if left > right {
//			return 1 + left
//		} else {
//			return 1 + right
//		}
//	}
//}

func isBalenced(node *TreeNode) bool {
	if node == nil {
		return true
	} else {
		return math.Abs(float64(maxDepth(node.Left))- float64(maxDepth(node.Right))) < 2 && isBalenced(node.Left) && isBalenced(node.Right)
	}
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}else {
		right := maxDepth(node.Right)
		left := maxDepth(node.Left)
		if right > left {
			return 1 + right
		} else {
			return 1 + left
		}
	}
}
