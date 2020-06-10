package main

import (
	. "github.com/isdamir/gotype"
)

// 链表逆序
func main() {
	head := &LNode{}
	CreateNode(head, 8)
	PrintNode("逆序前：", head)
}

// 就地逆序
func Reverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var pre *LNode
	var cur *LNode
	next := node.Next
	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
}

func RecursiveReverseChild(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	newHead := RecursiveReverseChild(node.Next)
	node.Next = node
	node.Next = nil
	return newHead
}

func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

func InsertReverse(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	var cur *LNode
	var next *LNode
	cur = node.Next.Next
	node.Next.Next = nil
	if cur != nil {
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}
