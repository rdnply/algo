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

func (l *list) findNthIter(n int) int {
	length := 0
	curr := l.head
	for curr != nil {
		length++
		curr = curr.next
	}

	if length < n {
		return -1
	}

	curr = l.head
	for i := 1; i < length-n+1; i++ {
		curr = curr.next
	}

	return curr.data
}

func (l *list) findNthWithTwoPointers(n int) int {
	if l.head == nil {
		return -1
	}

	main, temp := l.head, l.head

	for i := 0; i < n; i++ {
		if temp == nil {
			return -1
		}

		temp = temp.next
	}

	for temp != nil {
		temp = temp.next
		main = main.next
	}

	return main.data
}

func main() {
	l := NewList()

	l.push(5)
	l.push(4)
	l.push(3)
	l.push(2)
	l.push(1)

	fmt.Println(l.findNthIter(1))
	fmt.Println(l.findNthIter(2))

	fmt.Println(l.findNthWithTwoPointers(1))
	fmt.Println(l.findNthWithTwoPointers(2))
}
