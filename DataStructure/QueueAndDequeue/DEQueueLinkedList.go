package main

import "fmt"

type Node struct {
	element int
	next    *Node
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
	newest := &Node{e, nil}
	if d.isempty() {
		d.front = newest
		d.rear = newest
	} else {
		newest.next = d.front
		d.front = newest
	}
	d.size++
}

func (d *DEQue) addlast(e int) {
	newest := &Node{e, nil}
	if d.isempty() {
		d.front = newest
		d.rear = newest
	} else {
		d.rear.next = newest
		d.rear = newest
	}
	d.size++
}

func (d *DEQue) removefirst() {
	if d.isempty() {
		fmt.Println("DEQue is empty")
		return
	}
	e := d.front.element
	d.front = d.front.next
	d.size--
	fmt.Println(e)
}

func (d *DEQue) removelast() {
	if d.isempty() {
		fmt.Println("DEQue is empty")
		return
	} else {
		p := d.front
		i := 1
		for i < d.len()-1 {
			p = p.next
			i++
		}
		d.rear = p
		p = p.next
		e := p.element
		d.rear.next = nil
		d.size--
		fmt.Println(e)
	}

}

func (d *DEQue) first() {
	if d.isempty() {
		fmt.Println("DEQue is empty")
		return
	}
	fmt.Println(d.front.element)
}

func (d *DEQue) last() {
	if d.isempty() {
		fmt.Println("DEQue is empty")
		return
	}
	fmt.Println(d.rear.element)
}

func (d *DEQue) display() {
	p := d.front
	for p != nil {
		fmt.Print(p.element, "<-")
		p = p.next
	}
	fmt.Println()
}

func main() {
	d := &DEQue{}
	d.addfirst(21)
	d.addfirst(22)
	d.addfirst(23)
	d.addfirst(24)
	d.display()
	fmt.Println("length:", d.len())
	fmt.Println("-------------------------------------------")
	d.addlast(31)
	d.addlast(32)
	d.addlast(33)
	d.addlast(34)
	d.display()
	fmt.Println("length:", d.len())
	fmt.Println("-------------------------------------------")
	d.removefirst()
	d.removelast()
	d.display()
	fmt.Println("length:", d.len())
	fmt.Println("-------------------------------------------")
	d.first()
	d.last()
	d.display()
	fmt.Println("length:", d.len())
	fmt.Println("-------------------------------------------")
}
