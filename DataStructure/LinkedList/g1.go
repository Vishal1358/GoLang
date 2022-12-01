package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Singly struct {
	head *Node
	tail *Node
	size int
}

func (s *Singly) len() int {
	return s.size
}

func (s *Singly) isempty() bool {
	return s.size == 0
}

func (s *Singly) addlast(e int) {
	new := &Node{e, nil}
	if s.isempty() {
		s.head = new
		s.tail = new
	} else {
		s.tail.next = new
		s.tail = new
	}
	s.size++
}

func (s *Singly) addany(e, position int) {
	new := &Node{e, nil}
	if s.isempty() {
		s.head = new
		s.tail = new
	}
	p := s.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	new.next = p.next
	p.next = new
	s.size++

}

func (s *Singly) addfirst(e int) {
	new := &Node{e, nil}
	if s.isempty() {
		s.head = new
		s.tail = new
	} else {
		new.next = s.head
		s.head = new
	}
	s.size++
}

func (s *Singly) removefirst() {
	if s.isempty() {
		fmt.Println("Array is empty")
		return
	}
	e := s.head.data
	s.head = s.head.next
	s.size--
	fmt.Println("removed element:", e)
}

func (s *Singly) display() {
	p := s.head
	for p != nil {
		fmt.Print(p.data, "-->")
		p = p.next
	}
	fmt.Println()
}

func main() {
	A := Singly{}
	A.addfirst(25)
	A.addfirst(26)
	A.addfirst(27)
	A.addfirst(28)
	A.display()
	fmt.Println("------------------------------------")
	A.addlast(35)
	A.addlast(36)
	A.addlast(37)
	A.display()
	A.addany(50, 5)
	A.display()
	fmt.Println("------------------------------------")
	A.addany(70, 7)
	A.display()
	fmt.Println("------------------------------------")
	A.removefirst()
	A.display()
	fmt.Println("------------------------------------")
}
