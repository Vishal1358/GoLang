package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	front *Node
	rear  *Node
	size  int
}

func (s *Stack) len() int {
	return s.size
}

func (s *Stack) isempty() bool {
	return s.size == 0
}

func (s *Stack) push(e int) {
	new := &Node{e, nil}
	if s.isempty() {
		s.front = new
		s.rear = new
	} else {
		new.next = s.front
		s.front = new
	}
	s.size++
}

func (s *Stack) pop() {
	if s.isempty() {
		fmt.Println("Stack is empty")
		return
	} else {
		e := s.front.data
		s.front = s.front.next
		// s.front = nil
		s.size--
		fmt.Println(e)

	}

}

func (s *Stack) top() int {
	if s.isempty() {
		fmt.Println("Stack is empty")
	}
	return s.front.data
}

func (s *Stack) display() {
	p := s.front
	i := 0
	for p != nil {
		fmt.Print(p.data, "->")
		p = p.next
		i++
	}
	fmt.Println()
}

func main() {
	A := &Stack{}
	A.push(21)
	A.push(22)
	A.push(23)
	A.push(24)
	A.display()
	fmt.Println("-----------------------------------------")
	A.pop()
	A.display()
	fmt.Println("-----------------------------------------")
	A.pop()
	A.display()
	fmt.Println("-----------------------------------------")
	fmt.Println(A.top())
	A.display()
	fmt.Println("-----------------------------------------")
}
