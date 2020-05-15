package main
/**
一、增加
---------------------
    1、  	依次输入：9,1,5，3,2,1,5,4,7,8，-6，-3，-9，-11
            输出：中序遍历输出，已经排序好的数据。（广搜输出）

二、删除
---------------------
    1、		输入：5
            输出：中序遍历输出，已经排序好的数据。（广搜输出）
    2、  	输入：-3
            输出：中序遍历输出，已经排序好的数据。（广搜输出）
三、左左情况旋转
---------------------
	1、		输入：	一棵（链式或顺序）树，内容：
					|-5(1)
					|---3(1)
					|-----2(1)
					|-------1(1)
					|-----4(1)
					|---6(1)
					--------------------
			输出：	|-3(1)
					|---2(1)
					|-----1(1)
					|---5(1)
					|-----4(1)
					|-----6(1)
 */

func toTestLL(args []int) *avlTree {
	rt := new(avlTree)

	for _, v := range args {
		rt.insert(&rt.root, v)
	}

	return rt
}

func (ctx *avlTree)insert(node **avlTreeNode,data int){
	if *node == nil {
		*node = new(avlTreeNode)
		(*node).data = data
		return
	}

	if (*node).data > data {
		ctx.insert(&(*node).left,data)
	} else if (*node).data < data{
		ctx.insert(&(*node).right,data)
	} else {
		(*node).freq++
	}

	(*node).height = ctx.max(ctx.getNodeHeight((*node).left), ctx.getNodeHeight((*node).right)) + 1
}
