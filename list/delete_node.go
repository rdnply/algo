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

func (l *list) delete(k int) {
	var curr, prev *node
	curr = l.head

	if curr != nil && curr.data == k {
		l.head = curr.next
		return
	}

	for curr != nil && curr.data != k {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		return
	}

	prev.next = curr.next
}

func (l *list) deletePos(pos int) {
	if l.head == nil {
		return
	}

	if pos == 0 {
		l.head = l.head.next
		return
	}

	var curr *node
	curr = l.head

	for i := 0; curr != nil && i < pos-1; i++ {
		curr = curr.next
	}

	if curr == nil || curr.next == nil {
		return
	}

	next := curr.next.next

	curr.next = next
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
	l.push(7)
	l.push(1)
	l.push(3)
	l.push(2)

	// l.print()
	// l.delete(1)
	// l.print()
	// l.delete(7)
	// l.print()

	l.print()
	for _, v := range []int{0, 2, 7, 1} {
		l.deletePos(v)
		l.print()
	}

}
