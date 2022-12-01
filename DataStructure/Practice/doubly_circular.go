package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyCircular struct {
	head *Node
	tail *Node
	size int
}

func (d *DoublyCircular) len() int {
	return d.size
}

func (d *DoublyCircular) isempty() bool {
	return d.size == 0
}

func (d *DoublyCircular) addfirst(e int) {
	new := &Node{e, nil, nil}
	if d.isempty() {
		d.head = new
		d.tail = new
	}
	new.next = d.head
	d.head.prev = new
	d.tail.next = new
	new.prev = d.tail
	d.head = new
	d.size++
}

func (d *DoublyCircular) addlast(e int) {
	new := &Node{e, nil, nil}
	if d.isempty() {
		d.head = new
		d.tail = new
	} else {
		d.tail.next = new
		new.prev = d.tail
		new.next = d.head
		d.head.prev = new
		d.tail = new
	}
	d.size++
}

func (d *DoublyCircular) addany(e, position int) {
	if position == 1 {
		d.addfirst(e)
	} else if position >= d.len() {
		d.addlast(e)
	} else {
		newest := &Node{e, nil, nil}
		p := d.head
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		newest.prev = p
		newest.next = p.next
		p.next.prev = newest
		p.next = newest
		d.size++
	}
}

func (d *DoublyCircular) removefirst() {
	if d.isempty() {
		fmt.Println("Array is empty")
		return
	} else {
		e := d.head.data
		d.tail.next = d.head.next
		d.head.prev = d.tail
		d.head = d.head.next
		d.size--
		fmt.Println(e)
	}
}

func (d *DoublyCircular) removelast() {
	if d.isempty() {
		fmt.Println("Array is empty")
		return
	} else {
		e := d.tail.data
		d.head.prev = d.tail.prev
		d.tail.prev.next = d.head
		d.tail = d.tail.prev
		d.size--
		fmt.Println(e)
	}
}

func (d *DoublyCircular) removeany(position int) {
	if d.isempty() {
		fmt.Println("Array is empty")
		return
	} else if position == 1 {
		d.removefirst()
	} else if position >= d.len() {
		d.removelast()
	} else {
		p := d.head
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		e := p.next.data
		p.next = p.next.next
		p.next.next.prev = p
		d.size--
		fmt.Println(e)
	}

}

func (d *DoublyCircular) display() {
	p := d.head
	i := 0
	for i < d.len() {
		fmt.Print(p.data, "->")
		p = p.next
		i++
	}
	fmt.Println()
}

func (d *DoublyCircular) displayprev() {
	p := d.tail
	i := 0
	for i < d.len() {
		fmt.Print(p.data, "<-")
		p = p.prev
		i++
	}
	fmt.Println()
}

func main() {
	A := &DoublyCircular{}
	A.addfirst(25)
	A.addfirst(26)
	A.addfirst(27)
	A.addfirst(28)
	A.display()
	A.displayprev()
	fmt.Println("----------------------------------------")
	A.addlast(30)
	A.addlast(31)
	A.addlast(32)
	A.addlast(33)
	A.display()
	A.displayprev()
	fmt.Println("----------------------------------------")
	A.addany(50, 5)
	A.display()
	A.displayprev()
	fmt.Println("----------------------------------------")
	A.addany(100, 10)
	A.display()
	A.displayprev()
	fmt.Println("----------------------------------------")
	A.removefirst()
	A.display()
	A.displayprev()
	fmt.Println("----------------------------------------")
	A.removelast()
	A.display()
	A.displayprev()
	fmt.Println("----------------------------------------")
	A.removeany(8)
	A.display()
	A.displayprev()
	fmt.Println("----------------------------------------")
}
