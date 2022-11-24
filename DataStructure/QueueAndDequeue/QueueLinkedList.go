package main

import "fmt"

type node struct {
	element int
	next    *node
}

type QueueLinkedList struct {
	front *node
	rear  *node
	size  int
}

func (q *QueueLinkedList) len() int {
	return q.size
}

func (q *QueueLinkedList) isempty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) enqueue(e int) {
	newest := &node{e, nil}
	if q.isempty() {
		q.front = newest
	} else {
		q.rear.next = newest
	}
	q.rear = newest
	q.size++
}

func (q *QueueLinkedList) deque() {
	if q.isempty() {
		fmt.Println("Queue is empty")
		return
	}
	e := q.front.element
	q.front = q.front.next
	q.size--
	fmt.Println(e)
}

func (q *QueueLinkedList) first() int {
	if q.isempty() {
		fmt.Println("Queue is empty")
		return 0
	}
	return q.front.element
}

func (q *QueueLinkedList) display() {
	p := q.front
	for p != nil {
		fmt.Print(p.element, "<--")
		p = p.next
	}
	fmt.Println()
}

func main() {
	q := QueueLinkedList{}
	q.enqueue(21)
	q.enqueue(22)
	q.enqueue(23)
	q.enqueue(24)
	q.enqueue(25)
	q.display()
	fmt.Println("len: ", q.len())
	fmt.Println("<---------------------------------------")
	q.deque()
	q.display()
	fmt.Println("len: ", q.len())
	fmt.Println("<---------------------------------------")
	q.display()
	fmt.Println("first element: ", q.first())
	fmt.Println("len: ", q.len())
	fmt.Println("<---------------------------------------")

}
