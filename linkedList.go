package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func linkedList() {
	l := &List{}
	l.Push(23)
	l.Push(42)
	l.Push(1)
	l.Push(43)

	n := l.First()
	fmt.Println("forward iteration")
	for {
		fmt.Println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}

	n = l.Last()
	fmt.Println("reverse iteration")
	for {
		fmt.Println(n.value)
		n = n.Prev()
		if n == nil {
			break
		}
	}
}
