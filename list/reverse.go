package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head *node
}

func NewList() *list {
	return &list{}
}

func (l *list) push(k int) {
	newNode := &node{k, nil}
	if l.head == nil {
		l.head = newNode
		return
	}

	newNode.next = l.head
	l.head = newNode
}

func (l list) print() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}
	fmt.Println()
}

func (l *list) reverse() {
	var prev, curr, next *node
	curr = l.head

	for curr != nil {
		next = curr.next
		curr.next = prev

		prev = curr
		curr = next
	}

	l.head = prev
}

func reverse(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}

	rest := reverse(head.next)

	head.next.next = head
	head.next = nil

	return rest
}

func main() {
	l := NewList()
	l.push(3)
	l.push(2)
	l.push(1)

	l.print()
	res := reverse(l.head)
	// l.reverse()
	for res != nil {
		fmt.Print(res.data, " ")
		res = res.next
	}
	// l.print()
}
