package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type DEQue struct {
	front *Node
	rear  *Node
	size  int
}

func (d *DEQue) len() int {
	return d.size
}
func (d *DEQue) isempty() bool {
	return d.size == 0
}
func (d *DEQue) addfirst(e int) {
	new := &Node{e, nil}
	if d.isempty() {
		d.front = new
		d.rear = new
	}
	new.next = d.front
	d.front = new
	d.size++

}
func (d *DEQue) addlast(e int) {
	new := &Node{e, nil}
	if d.isempty() {
		d.front = new
		d.rear = new
	}
	d.rear.next = new
	d.rear = new
	d.size++
}
func (d *DEQue) removefirst() {
	if d.isempty() {
		fmt.Println("DEQue is empty")
		return
	}
	e := d.front.data
	d.front = d.front.next
	d.size--
	fmt.Println(e)
}
func (d *DEQue) removelast() {
	if d.isempty() {
		fmt.Println("DEQue is empty")
		return
	}
	p := d.front
	i := 1
	for i < d.len()-1 {
		p = p.next
		i++
	}
	e := p.next.data
	d.rear = p
	p.next = nil
	d.size--
	fmt.Println(e)
}
func (d *DEQue) top() int {
	if d.isempty() {
		fmt.Println("DEQue is empty")
	}
	return d.front.data
}
func (d *DEQue) bottom() int {
	if d.isempty() {
		fmt.Println("DEQue is empty")
	}
	return d.rear.data
}

func (d *DEQue) display() {
	p := d.front
	i := 0
	for i < d.len() {
		fmt.Print(p.data, "->")
		p = p.next
		i++
	}
	fmt.Println()
}

func main() {
	A := &DEQue{}
	A.addfirst(25)
	A.addfirst(26)
	A.addfirst(27)
	A.addfirst(28)
	A.display()
	fmt.Println("-------------------------------")
	A.addlast(30)
	A.addlast(31)
	A.addlast(32)
	A.addlast(33)
	A.display()
	fmt.Println("-------------------------------")
	A.removefirst()
	A.removelast()
	A.display()
	fmt.Println("-------------------------------")
	fmt.Println(A.top(), A.bottom())
}
