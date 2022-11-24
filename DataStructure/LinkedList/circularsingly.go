package main

import "fmt"

type Node struct {
	element int
	next    *Node
}

type CircularLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *CircularLinkedList) len() int {
	return l.size
}

func (l *CircularLinkedList) isempty() bool {
	return l.size == 0
}

func (l *CircularLinkedList) addfirst(e int) {
	new := &Node{e, nil}
	if l.isempty() {
		new.next = new
		l.tail = new

	} else {
		new.next = l.head
		// l.tail.next = new
	}
	l.head = new
	l.size++
}

func (l *CircularLinkedList) addlast(e int) {
	new := &Node{e, nil}
	if l.isempty() {
		l.head = new
		new.next = new
	} else {
		new.next = l.head
		l.tail.next = new
	}
	l.tail = new
	l.size++
}

func (l *CircularLinkedList) addany(e, position int) {
	new := &Node{e, nil}
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	new.next = p.next
	p.next = new
	l.size++
}

func (l *CircularLinkedList) removefirst() {
	if l.isempty() {
		fmt.Println("list is empty")
		return
	} else {
		l.head = l.head.next
		l.tail.next = l.head
		l.size--
	}
	if l.isempty() {
		l.head = nil
		l.tail = nil
	}
}

func (l *CircularLinkedList) removelast() {
	if l.isempty() {
		fmt.Println("list is empty")
	} else {
		p := l.head
		i := 1
		for i < l.len()-1 {
			p = p.next
			i++
		}
		p.next = p.next.next
		l.size--
	}
}

func (l *CircularLinkedList) removeany(position int) {
	if l.isempty() {
		fmt.Println("list is empty")
	} else {
		p := l.head
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		p.next = p.next.next
		l.size--
	}
}

func (l *CircularLinkedList) display() {
	p := l.head
	i := 0
	for i < l.len() {
		fmt.Print(p.element, "->")
		p = p.next
		i++
	}
	fmt.Println()
}

func main() {
	O := &CircularLinkedList{}
	O.addfirst(25)
	O.addfirst(26)
	O.addfirst(27)
	O.addfirst(28)
	O.addfirst(29)
	O.display()
	fmt.Println()
	O.addlast(31)
	O.addlast(32)
	O.addlast(33)
	O.addlast(34)
	O.display()
	fmt.Println("----------------------------------------------------")
	O.addany(100, 3)
	O.display()
	fmt.Println("----------------------------------------------------")
	O.removefirst()
	O.display()
	fmt.Println("----------------------------------------------------")
	O.removelast()
	O.display()
	fmt.Println("----------------------------------------------------")
	O.removeany(2)
	O.display()
	fmt.Println("----------------------------------------------------")
}
