package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyCircular struct {
	head *Node
	tail *Node
	size int
}

func (s *SinglyCircular) len() int {
	return s.size
}

func (s *SinglyCircular) isempty() bool {
	return s.size == 0
}

func (s *SinglyCircular) addfirst(e int) {
	new := &Node{e, nil}
	if s.isempty() {
		s.head = new
		s.tail = new
	} else {
		s.tail.next = new
		new.next = s.head
		s.head = new
	}
	s.size++
}

func (s *SinglyCircular) addlast(e int) {
	new := &Node{e, nil}
	if s.isempty() {
		s.head = new
		s.tail = new
	} else {
		new.next = s.head
		s.tail.next = new
		s.tail = new
	}
	s.size++
}

func (s *SinglyCircular) addany(e int, position int) {
	new := &Node{e, nil}
	if s.isempty() {
		s.head = new
		s.tail = new
	} else {
		p := s.head
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		new.next = p.next
		p.next = new
	}
	s.size++
}

func (s *SinglyCircular) removefirst() {
	if s.isempty() {
		fmt.Print("Array is empty")
		return
	}
	e := s.head.data
	s.head = s.head.next
	s.size--
	fmt.Println(e)
}
func (s *SinglyCircular) removelast() {
	if s.isempty() {
		fmt.Print("Array is empty")
		return
	}
	p := s.head
	i := 1
	for i < s.len()-1 {
		p = p.next
		i++
	}
	e := p.next.data
	p.next = p.next.next
	s.size--
	fmt.Println(e)

}

func (s *SinglyCircular) removeany(position int) {
	if s.isempty() {
		fmt.Println("Array is empty")
		return
	}
	p := s.head
	i := 1
	for i < position-1 {
		p = p.next
		i++
	}
	e := p.next.data
	p.next = p.next.next
	s.size--
	fmt.Println(e)
}

func (s *SinglyCircular) display() {
	p := s.head
	i := 0
	for i < s.len() {
		fmt.Print(p.data, "->")
		p = p.next
		i++
	}
	fmt.Println()
}

func main() {
	A := &SinglyCircular{}
	A.addfirst(25)
	A.addfirst(26)
	A.addfirst(27)
	A.addfirst(28)
	A.display()
	fmt.Println("-------------------------------------")
	A.addlast(30)
	A.addlast(31)
	A.addlast(32)
	A.addlast(33)
	A.display()
	fmt.Println("-------------------------------------")
	A.addany(50, 5)
	A.display()
	fmt.Println("-------------------------------------")
	A.addany(90, 10)
	A.display()
	fmt.Println("-------------------------------------")
	A.removefirst()
	A.display()
	fmt.Println("-------------------------------------")
	A.removelast()
	A.display()
	fmt.Println("-------------------------------------")
	A.removeany(3)
	A.display()
	fmt.Println("-------------------------------------")
	A.removeany(3)
	A.display()
	fmt.Println("-------------------------------------")
}
