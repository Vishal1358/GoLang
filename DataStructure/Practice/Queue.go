package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

func (q *Queue) len() int {
	return q.size
}

func (q *Queue) isempty() bool {
	return q.size == 0
}

func (q *Queue) enqueue(e int) {
	new := &Node{e, nil}
	if q.isempty() {
		q.front = new
		q.rear = new
	}
	new.next = q.front
	q.front = new
	q.size++
}

func (q *Queue) dqueue() {
	if q.isempty() {
		fmt.Println("Queue is empty")
		return
	}
	e := q.front.data
	q.front = q.front.next
	q.size--
	fmt.Println(e)

}

func (q *Queue) top() int {
	if q.isempty() {
		fmt.Println("Queue is empty")
	}
	return q.front.data
}

func (q *Queue) display() {
	p := q.front
	i := 0
	for i < q.len() {
		fmt.Print(p.data, "->")
		p = p.next
		i++
	}
	fmt.Println()
}

func main() {
	A := &Queue{}
	A.enqueue(25)
	A.enqueue(26)
	A.enqueue(27)
	A.enqueue(28)
	A.display()
	fmt.Println("-------------------------------------")
	A.dqueue()
	// A.dqueue()
	// A.dqueue()
	// A.dqueue()
	// A.dqueue()
	// A.display()
	fmt.Println("-------------------------------------")
	fmt.Println(A.top())

}
