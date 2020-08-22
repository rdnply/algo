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

func (l list) findCycle() bool {
	slow, fast := l.head, l.head

	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	l := NewList()
	l.push(4)
	l.push(3)
	l.push(2)
	l.push(1)

	l.head.next.next.next = l.head

	fmt.Println(l.findCycle())
}
