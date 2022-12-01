package main

import "fmt"

type Node struct {
	element int
	next    *Node
}

type StackLinkedList struct {
	_top *Node
	size int
}

func (l *StackLinkedList) len() int {
	return l.size
}

func (l *StackLinkedList) isempty() bool {
	return l.size == 0
}

func (l *StackLinkedList) push(e int) {
	newest := &Node{e, nil}
	if l.isempty() {
		l._top = newest
	} else {
		newest.next = l._top
		l._top = newest
	}
	l.size++
}

func (l *StackLinkedList) pop() {
	if l.isempty() {
		fmt.Println("Stack is empty")
		return
	}
	e := l._top.element
	l._top = l._top.next
	l.size--
	fmt.Println(e)
}

func (l *StackLinkedList) top() int {
	if l.isempty() {
		fmt.Println("Stack is empty")
	}
	return l._top.element
}

func (l *StackLinkedList) display() {
	p := l._top
	for p != nil {
		fmt.Print(p.element, "-->")
		p = p.next
	}
	fmt.Println()
}

func main() {
	S := StackLinkedList{}
	S.push(25)
	S.push(26)
	S.push(27)
	fmt.Println("length:", S.len())
	S.display()
	S.pop()
	S.display()
	fmt.Println("length:", S.len())

	fmt.Println(S.top())
	S.display()
	fmt.Println("length:", S.len())

}
