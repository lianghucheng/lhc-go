package main

import (
	"fmt"
	"math"
)
//高度平衡二叉树，左子树和右子树的高度相差不超过1，并且左子树和右子树也是高度平衡二叉树
type avlTreeNode struct {
	val int
	left *avlTreeNode
	right *avlTreeNode
}

func main() {
	fmt.Println(isBalance(new(avlTreeNode)))
}

func isBalance(node *avlTreeNode)bool {
	if node == nil {
		return true
	}
	return math.Abs(float64(depth(node.right) - depth(node.left))) <= 1 && isBalance(node.left) && isBalance(node.right)
}

func depth(node *avlTreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.left)
	right := depth(node.right)
	if left > right {
		return 1 + left
	} else {
		return 1 + right
	}
}

