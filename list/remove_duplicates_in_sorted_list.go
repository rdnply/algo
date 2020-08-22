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

func (l *list) removeDuplicates() {
	if l.head == nil {
		return
	}

	curr := l.head

	for curr.next != nil {
		if curr.data == curr.next.data {
			next := curr.next.next
			curr.next = next
		} else {
			curr = curr.next
		}
	}
}

func (l list) print() {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}

	fmt.Println()
}

func main() {
	l := NewList()

	l.push(5)
	l.push(5)
	l.push(5)
	l.push(5)
	l.push(5)

	l.print()
	l.removeDuplicates()
	l.print()
}
