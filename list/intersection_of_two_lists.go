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

func intersection(l1 *list, l2 *list) *list {
	tail := NewList()

	for l1.head != nil && l2.head != nil {
		if l1.head.data == l2.head.data {
			tail.push(l1.head.data)
			l1.head = l1.head.next
			l2.head = l2.head.next
		} else if l1.head.data < l2.head.data {
			l1.head = l1.head.next
		} else {
			l2.head = l2.head.next
		}
	}

	return tail
}

func intersectionRecur(a *node, b *node) *node {
	if a == nil || b == nil {
		return nil
	}

	if a.data < b.data {
		return intersectionRecur(a.next, b)
	}

	if a.data > b.data {
		return intersectionRecur(a, b.next)
	}

	temp := &node{a.data, nil}
	temp.next = intersectionRecur(a.next, b.next)

	return temp
}

func print(l list) {
	curr := l.head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}

	fmt.Println()
}

func printFromHead(head *node) {
	curr := head
	for curr != nil {
		fmt.Print(curr.data, " ")
		curr = curr.next
	}

	fmt.Println()
}

func main() {
	a := NewList()
	b := NewList()

	a.push(3)
	a.push(1)
	a.push(1)

	b.push(3)
	b.push(2)
	b.push(1)

	// res := intersection(a, b)
	// print(*res)

	fmt.Println(a.head)
	resRec := intersectionRecur(a.head, b.head)
	printFromHead(resRec)
}
