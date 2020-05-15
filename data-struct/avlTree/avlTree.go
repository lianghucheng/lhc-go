package main

import (
	"fmt"
)

type avlTree struct {
	root *avlTreeNode
}

type avlTreeNode struct {
	data	int
	left	*avlTreeNode
	right	*avlTreeNode
	height	int
	freq 	int
}

func (ctx *avlTree)getNodeHeight(node *avlTreeNode) int {
	if node == nil {
		return -1
	} else {
		return node.height
	}
}

func (ctx *avlTree)max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (ctx *avlTree)SingleRotateLeft(node **avlTreeNode) {
	rootPtr := node
	oldRoot := *rootPtr
	newRoot := oldRoot.left
	oldRoot.left = newRoot.right
	newRoot.right = oldRoot
	*rootPtr = newRoot
}

func (ctx *avlTree)SingleRotateRight(node **avlTreeNode) {
	rootPtr := node
	oldRoot := *rootPtr
	newRoot := oldRoot.right
	oldRoot.right = newRoot.left
	newRoot.left = oldRoot
	*rootPtr = newRoot
}

func (ctx *avlTree)DoubleRotateLR(node **avlTreeNode) {
	ctx.SingleRotateRight(&(*node).left)
	ctx.SingleRotateLeft(node)
}

func (ctx *avlTree)DoubleRotateRL(node **avlTreeNode) {
	ctx.SingleRotateLeft(&(*node).right)
	ctx.SingleRotateRight(node)
}

func (ctx *avlTree)show(node *avlTreeNode) {
	if node != nil {
		node.show(0)
	}
}

func (ctx *avlTree)show2() {
	if ctx.root == nil {
		return
	}

	queue := []*avlTreeNode{ctx.root}
	currheight
	for ;len(queue) > 0; {
		q := queue[0]

		queue = queue[1:]
		q.printForm2()
		if q.left != nil {
			queue = append(queue, q.left)
		}
		if q.right != nil {
			queue = append(queue, q.right)
		}
		fmt.Println()
	}
}

func (ctx *avlTreeNode)show(height int) {
	if ctx == nil {
		return
	}
	height++
	ctx.printForm(height)
	ctx.left.show(height)
	ctx.right.show(height)
}

func (ctx *avlTreeNode)printForm(height int) {
	fmt.Print("|-")
	for i := 0; i < height - 1; i++ {
		fmt.Print("--------------")
	}
	fmt.Printf("%v(freq:%v,heigh:%v)\n",ctx.data, ctx.freq, ctx.height)
}

func (ctx *avlTreeNode)printForm2() {
	fmt.Print("   ")
	for i:=0;i<ctx.height; i++ {
		fmt.Print("       ")
	}
	fmt.Printf("%v(freq:%v,heigh:%v)",ctx.data, ctx.freq, ctx.height)
}