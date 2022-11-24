package main

import "fmt"

type Node struct {
	element int
	next    *Node
	prev    *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *DoublyLinkedList) len() int {
	return l.size
}

func (l *DoublyLinkedList) isempty() bool {
	return l.size == 0
}

func (l *DoublyLinkedList) addfirst(e int) {
	new := &Node{e, nil, nil}
	if l.head == nil && l.tail == nil {
		l.head = new
		l.tail = new
	} else {
		new.next = l.head
		l.head.prev = new
		l.head = new
	}
	l.size++
}

func (l *DoublyLinkedList) addlast(e int) {
	new := &Node{e, nil, nil}
	if l.head == nil && l.tail == nil {
		l.head = new
		l.tail = new
	} else {
		l.tail.next = new
		new.prev = l.tail
		l.tail = new
	}
	l.size++
}

func (l *DoublyLinkedList) addany(e int, position int) {
	new := &Node{e, nil, nil}
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	new.next = p.next
	p.next.prev = new
	p.next = new
	new.prev = p
	l.size++

}

func (l *DoublyLinkedList) removefirst() {
	if l.isempty() {
		fmt.Println("list is empty")
		return
	} else {
		l.head = l.head.next
		l.head.prev = nil
		l.size--
	}
	if l.isempty() {
		l.head = nil
		l.tail = nil
	}
}

func (l *DoublyLinkedList) removelast() {
	if l.isempty() {
		fmt.Println("list is empty")
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
	}
	if l.isempty() {
		l.head = nil
		l.tail = nil
	}
}

func (l *DoublyLinkedList) removeany(position int) {
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	p.next = p.next.next
	p.next.prev = p
	l.size--
	if l.isempty() {
		l.head = nil
		l.tail = nil
	}
}

func (l *DoublyLinkedList) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.element, "-->")
		p = p.next
	}
	fmt.Println()
}

func (l *DoublyLinkedList) displayrev() {
	p := l.tail
	for p != nil {
		fmt.Print(p.element, "-->")
		p = p.prev
	}
	fmt.Println()
}

func main() {
	O := &DoublyLinkedList{}
	O.addfirst(25)
	O.addfirst(26)
	O.addfirst(27)
	O.addfirst(28)
	O.addfirst(29)
	O.display()
	O.displayrev()
	O.addlast(30)
	O.display()
	O.displayrev()
	O.addany(31, 6)
	O.display()
	O.displayrev()
	O.removefirst()
	O.display()
	O.displayrev()
	O.removelast()
	O.display()
	O.displayrev()
	O.removeany(3)
	O.display()
	O.displayrev()
}
