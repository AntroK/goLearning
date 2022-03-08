package main

import "fmt"

// 实现1-> 2 -> 3
// 	  2-> 3 -> 4
// 相加3-> 5 -> 7
// 指针要初始化！！！
type Nd struct {
	Value 	int
	Next	*Nd
	Pre		*Nd
}

type LinkList struct {
	root 	*Nd
	size 	int
	tail 	*Nd
}

func(l *LinkList) AddNode(node *Nd){
	if l.size == 0{
		l.root.Next = node
		l.tail = l.root
		l.size = 1
		return
	}
	node.Next = l.tail.Next //= l.root.Next
	l.tail.Next = node
	l.tail = node
	l.size += 1
}

func NewLinkList() *LinkList{
	return &LinkList{
		root: &Nd{},
		tail: &Nd{},
		size: 0,
	}
}

func add(l1, l2 *LinkList) *LinkList{
	var res = NewLinkList()
	if l1.size == 0 {
		return l2
	}
	if l2.size == 0{
		return l1
	}

	node1 := l1.root.Next
	node2 := l2.root.Next
	flag := 0
	for node1 != nil || node2 != nil {
		val := node1.Value  + node2.Value + flag
		if val > 9 {
			val = val - 10
			flag = 1
		}else{
			flag = 0
		}
		res.AddNode(&Nd{Value: val})
		node1 = node1.Next
		node2 = node2.Next
	}
	return res
}

func main(){
	node1 := &Nd{Value: 1}
	node2 := &Nd{Value: 2}
	node3 := &Nd{Value: 3}
	node4 := &Nd{Value: 4}

	ls1 := NewLinkList()
	ls2 := NewLinkList()
	ls1.AddNode(node1)
	ls1.AddNode(node2)
	ls2.AddNode(node3)
	ls2.AddNode(node4)
	res := add(ls1, ls2)
	fmt.Print(res)
}