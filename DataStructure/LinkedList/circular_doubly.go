package main

import "fmt"

type node struct {
	element int
	next    *node
	prev    *node
}

type CircularDoubly struct {
	haed *node
	tail *node
	size int
}

func (c *CircularDoubly) len() int {
	return c.size
}

func (c *CircularDoubly) isempty() bool {
	return c.size == 0
}

func (c *CircularDoubly) addfirst(e int) {
	newest := &node{e, nil, nil}
	if c.isempty() {
		c.haed = newest
		c.tail = newest
		newest.next = newest
		newest.prev = newest
	} else {
		newest.next = c.haed
		c.haed.prev = newest
		c.tail.next = newest
		c.haed = newest
	}
	c.size++
}

func (c *CircularDoubly) addlast(e int) {
	newest := &node{e, nil, nil}
	if c.isempty() {
		c.haed = newest
		c.tail = newest
		newest.next = newest
		newest.prev = newest
	} else {
		c.tail.next = newest
		newest.prev = c.tail
		c.tail = newest
		c.tail.next = c.haed
		c.haed.prev = c.tail
	}
	c.size++
}

func (c *CircularDoubly) display() {
	p := c.haed
	i := 0
	for i < c.len() {
		fmt.Print(p.element, "->")
		p = p.next
		i++
	}
	fmt.Println()
}

func (c *CircularDoubly) addany(e, position int) {
	if position == 1 {
		c.addfirst(e)
	} else if position >= c.len() {
		c.addlast(e)
	} else {
		newest := &node{e, nil, nil}
		p := c.haed
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		newest.prev = p
		newest.next = p.next
		p.next.prev = newest
		p.next = newest
		c.size++
	}

}

func (c *CircularDoubly) removefirst() {
	if c.isempty() {
		fmt.Println("List is empty")
		return
	}
	e := c.haed.element
	c.haed = c.haed.next
	c.haed.prev = c.tail
	c.size--
	fmt.Println("Removed element: ", e)
}

func (c *CircularDoubly) removelast() {
	if c.isempty() {
		fmt.Println("List is empty")
		return
	}
	e := c.tail.element
	c.tail = c.tail.prev
	c.tail.next = c.haed
	c.size--
	// return
	fmt.Println("Removed element: ", e)

}

func (c *CircularDoubly) removeany(position int) {
	if c.isempty() {
		fmt.Println("Array is empty")
		return
	} else if position == 1 {
		c.removefirst()
	} else if position >= c.len() {
		c.removelast()
	} else {
		p := c.haed
		i := 1
		for i < position-1 {
			p = p.next
			i++
		}
		e := p.next.element
		p.next = p.next.next
		p.next.next.prev = p
		c.size--
		fmt.Println(e)
	}
}

func (c *CircularDoubly) displayrev() {
	p := c.tail
	i := 0
	for i < c.len() {
		fmt.Print(p.element, "->")
		p = p.prev
		i++
	}
	fmt.Println()
}

func main() {
	l := &CircularDoubly{}
	l.addfirst(21)
	l.addfirst(22)
	l.addfirst(23)
	l.addfirst(24)
	l.display()
	l.displayrev()
	fmt.Println("Length :", l.len())
	fmt.Println("-----------------------------------------------------------------")
	l.addlast(30)
	l.addlast(31)
	l.display()
	l.displayrev()
	fmt.Println("Length :", l.len())
	fmt.Println("-----------------------------------------------------------------")
	l.addany(50, 5)
	l.display()
	l.displayrev()
	fmt.Println("Length :", l.len())
	fmt.Println("-----------------------------------------------------------------")
	l.addany(55, 8)
	l.display()
	l.displayrev()
	fmt.Println("Length :", l.len())
	fmt.Println("-----------------------------------------------------------------")
	// l.removefirst()
	// l.display()
	// l.displayrev()
	// fmt.Println("Length :", l.len())
	// fmt.Println("-----------------------------------------------------------------")
	// l.removelast()
	// l.display()
	// l.displayrev()
	// fmt.Println("Length :", l.len())
	// fmt.Println("-----------------------------------------------------------------")
	// l.removeany(2)
	// l.display()
	// l.displayrev()
	// fmt.Println("Length :", l.len())
	// fmt.Println("-----------------------------------------------------------------")

}
