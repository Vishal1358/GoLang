package main

import "fmt"

type Node struct {
	element int
	next    *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *SinglyLinkedList) len() int {
	return l.size
}

func (l *SinglyLinkedList) isempty() bool {
	return l.size == 0
}

func (l *SinglyLinkedList) addlast(e int) {
	new := &Node{e, nil}
	if l.head == nil {
		l.head = new
	} else {
		l.tail.next = new
	}
	l.tail = new
	l.size++

}

func (l *SinglyLinkedList) addfirst(e int) {
	new := &Node{e, nil}
	if l.head == nil {
		l.head = new
		l.tail = new
	} else {
		new.next = l.head
		l.head = new
	}
	l.size++
}

func (l *SinglyLinkedList) addany(e, position int) {
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

func (l *SinglyLinkedList) removefirst() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		p := l.head.next
		l.head = p
		l.size++
	}
	if l.head == nil {
		l.tail = nil
	}
}

func (l *SinglyLinkedList) removeend() {
	if l.isempty() {
		fmt.Println("list is empty")
	} else {
		p := l.head
		i := 1
		for i < l.len()-1 {
			p = p.next
			i++
		}
		p.next = nil
		l.size--
	}
}

func (l *SinglyLinkedList) removeany(position int) {
	p := l.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	p.next = p.next.next
	l.size--
}

func (l *SinglyLinkedList) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.element, "-->")
		p = p.next
	}
}

func main() {
	O := &SinglyLinkedList{}
	O.addlast(25)
	O.addlast(26)
	O.addlast(27)
	O.addlast(28)
	O.addlast(29)
	O.display()
	fmt.Println()
	O.addfirst(30)
	O.display()
	fmt.Println()
	O.addany(32, 3)
	O.display()
	fmt.Println()
	O.removefirst()
	O.display()
	fmt.Println()
	O.removeany(3)
	O.display()
	fmt.Println()
	O.removeend()
	O.display()
	fmt.Println()

}
