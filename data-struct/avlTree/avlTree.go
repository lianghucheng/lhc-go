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
	oldRoot.height = ctx.max(ctx.getNodeHeight(oldRoot.left), ctx.getNodeHeight(oldRoot.right))+ 1
	newRoot.height = ctx.max(ctx.getNodeHeight(newRoot.left), oldRoot.height) + 1
	*rootPtr = newRoot
}

func (ctx *avlTree)SingleRotateRight(node **avlTreeNode) {
	rootPtr := node
	oldRoot := *rootPtr
	newRoot := oldRoot.right
	oldRoot.right = newRoot.left
	newRoot.left = oldRoot
	oldRoot.height = ctx.max(ctx.getNodeHeight(oldRoot.left), ctx.getNodeHeight(oldRoot.right)) + 1
	newRoot.height = ctx.max(oldRoot.height, ctx.getNodeHeight(newRoot.right)) + 1
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

func (ctx *avlTree)insert(node **avlTreeNode, data int) {
	if *node == nil {
		*node = new(avlTreeNode)
		(*node).data = data
		return
	}

	if (*node).data > data {
		ctx.insert(&(*node).left,data)
		if ctx.getNodeHeight((*node).left) - ctx.getNodeHeight((*node).right) == 2 {
			if data < (*node).left.data {
				ctx.SingleRotateLeft(node)
			} else {
				ctx.DoubleRotateLR(node)
			}
		}
	} else if (*node).data < data{
		ctx.insert(&(*node).right,data)
		if ctx.getNodeHeight((*node).right)- ctx.getNodeHeight((*node).left) == 2 {
			if data > (*node).right.data {
				ctx.SingleRotateRight(node)
			} else {
				ctx.DoubleRotateRL(node)
			}
		}
	} else {
		(*node).freq++
	}

	(*node).height = ctx.max(ctx.getNodeHeight((*node).left), ctx.getNodeHeight((*node).right)) + 1
}

func (ctx *avlTree)insert2(node **avlTreeNode, data int) {
	if *node == nil {
		*node = new(avlTreeNode)
		(*node).data = data
		return
	}

	if (*node).data > data {
		ctx.insert2(&(*node).left,data)
		if ctx.getNodeHeight((*node).left) - ctx.getNodeHeight((*node).right) == 2 {
			if data < (*node).left.data {
				ctx.SingleRotateLeft(node)
			} else {
				ctx.DoubleRotateLR(node)
			}
		}
	} else if (*node).data < data{
		ctx.insert2(&(*node).right,data)
		if ctx.getNodeHeight((*node).right)- ctx.getNodeHeight((*node).left) == 2 {
			if data > (*node).right.data {
				ctx.SingleRotateRight(node)
			} else {
				ctx.DoubleRotateRL(node)
			}
		}
	} else {
		(*node).freq++
	}

	(*node).height = ctx.max(ctx.getNodeHeight((*node).left), ctx.getNodeHeight((*node).right)) + 1
}

func (ctx *avlTree)Insert(data int) {
	ctx.insert(&ctx.root, data)
}

func (ctx *avlTree)find(node *avlTreeNode, data int)*avlTreeNode{
	if node == nil {
		return nil
	}
	if node.data > data {
		return ctx.find(node.left, data)
	} else if node.data < data {
		return ctx.find(node.right, data)
	} else {
		return node
	}
}

func (ctx *avlTree)Find(data int) *avlTreeNode {
	return ctx.find(ctx.root, data)
}

func (ctx *avlTree)delete(node **avlTreeNode, data int) {
	if *node == nil {
		return
	}
	if data < (*node).data {
		ctx.delete(&(*node).left, data)
		if ctx.getNodeHeight((*node).right) - ctx.getNodeHeight((*node).left) == 2{
			if (*node).right.left != nil && ctx.getNodeHeight((*node).right.left) > ctx.getNodeHeight((*node).right.right) {
				ctx.DoubleRotateRL(node)
			} else {
				ctx.SingleRotateRight(node)
			}
		}
	} else if data > (*node).data {
		ctx.delete(&(*node).right, data)
		if ctx.getNodeHeight((*node).left) - ctx.getNodeHeight((*node).right) == 2{
			if (*node).left.right!=nil&&ctx.getNodeHeight((*node).left.right) > ctx.getNodeHeight((*node).left.left) {
				ctx.DoubleRotateLR(node)
			} else {
				ctx.SingleRotateLeft(node)
			}
		}
	} else {
		if (*node).left !=nil && (*node).right != nil {
			temp := (*node).right
			for temp.left != nil {
				temp = temp.left
			}
			(*node).data = temp.data
			(*node).freq = temp.freq
			ctx.delete(&(*node).right, temp.data)
			if ctx.getNodeHeight((*node).left) - ctx.getNodeHeight((*node).right) ==2 {
				if (*node).left.right!=nil && ctx.getNodeHeight((*node).left.right) > ctx.getNodeHeight((*node).left.left) {
					ctx.DoubleRotateLR(node)
				} else {
					ctx.SingleRotateLeft(node)
				}
			}
		} else {
			temp := *node
			if (*node).left == nil {
				*node = (*node).right
			} else if (*node).right == nil {
				*node = (*node).left
			}
			temp = nil
			_ = temp
		}
	}
	if *node == nil {
		return
	}
	(*node).height = ctx.max(ctx.getNodeHeight((*node).left), ctx.getNodeHeight((*node).right))+1
}

func (ctx *avlTree)Delete(data int) {
	ctx.delete(&ctx.root, data)
}

func (ctx *avlTree)insubTree(node *avlTreeNode) {
	if node == nil {
		return
	}
	ctx.insubTree(node.left)
	fmt.Print(node.data," ")
	ctx.insubTree(node.right)
}

func (ctx *avlTree)InsubTree() {
	ctx.insubTree(ctx.root)
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
	for ;len(queue) > 0; {
		tempQueue := []*avlTreeNode{}
		for _, q := range queue {
			q.printForm2()
			if q.left != nil {
				tempQueue = append(tempQueue, q.left)
			}
			if q.right != nil {
				tempQueue = append(tempQueue, q.right)
			}
		}

		queue = tempQueue

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
	fmt.Print("                 ")
	for i:=0;i<ctx.height; i++ {
		fmt.Print("         ")
	}
	fmt.Printf("%v(freq:%v,heigh:%v)",ctx.data, ctx.freq, ctx.height)
}