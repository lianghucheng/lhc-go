package llist

import "fmt"

//终止原因：	都是自己总结的，太耗时间了。决定直接吸收前人的经验
//当前小结：	遍历链表时：
// 				1、要直观。意思是遍历到哪一个就抓住当前那一个节点，不要抓前一个节点或者后一个节点。
//				2、计数器count更新的值，要是当前已到节点之前的总长度，不多一也不少一。
//				3、链表带头，永远不动，直至销毁
//				4、count-2为当前真正索引。
//				5、所有操作都在节点后移之前。
//				6、取出指定节点，需要用pre记住，不能用后移过的节点，因为会是空指针。
//				7、遍历节点是否做成一个函数，将应该进行的操作用函数传参。
//时间：2020-05-11
func main() {
	mp1 := make(map[string]int)
	var mp2 map[string]int

	fmt.Println(mp1 == nil, mp2 == nil)
}

type singleChain struct {
	size int
	head *singleChainNode
}

type singleChainNode struct {
	data int
	next *singleChainNode
}

func createSingleChain() (rt *singleChain) {
	rt = new(singleChain)
	rt.head = new(singleChainNode)
	return
}

func (ctx *singleChain) destroySingleChain() {

}

func (ctx *singleChain) add (data, index int) {
	head := ctx.head
	count := 0

	var pre *singleChainNode
	for ;head != nil && index > 0; {
		pre = head
		head = head.next
		count++
		if count - 2 == index - 1 { //count减2是当前节点的索引，index是将要加到的位置，index - 1是将要加到的位置的前驱节点
			break//找到前驱节点结束
		}
	}
	//如果index大于了链表的长度，那么就追加到最后一个节点。此处不需要多余的操作，当前程序默认支持这种情况。只是在此说明一下特殊情况下会发生的事情。
	//如果index小于零，则添加到索引0的位置，无需遍历链表。此处不需要多余的操作，当前程序默认支持这种情况。只是在此说明一下特殊情况下会发生的事情。

	node := new(singleChainNode)
	node.data = data
	node.next = pre.next
	pre.next = node

	//更新长度
	ctx.size++
}

func (ctx *singleChain) delete(index int) {

}

func (ctx *singleChain) get(index int) (rt int) {
	return
}

func (ctx *singleChain) set(data, index int) {

}

func (ctx *singleChain) show() {
	head := ctx.head
	count := 0
	for ;head != nil; {
		count++
		if count - 2 >= 0 {
			fmt.Print(head.data,"  ")
		}
		head = head.next
	}
	fmt.Println()
}


