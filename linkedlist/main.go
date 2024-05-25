package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(value int) {
	if l.Head == nil {
		l.Head = &Node{Value: value}
		return
	}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = &Node{Value: value}
}

func (l *LinkedList) Delete(value int) {
	if l.Head == nil {
		return
	}

	current := l.Head

	if current.Value == value {
		l.Head = current.Next
		return
	}

	for current.Next != nil {
		if value == current.Next.Value {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

func (l *LinkedList) Get(index int) *Node {
	if l.Head == nil {
		return nil
	}

	current := l.Head
	counter := 0

	for current != nil {
		if counter == index {
			return current
		}
		counter += 1
		current = current.Next
	}

	return nil
}

func (l *LinkedList) Reverse() {
	if l.Head == nil {
		return
	}

	var previus *Node
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = previus
		previus = current
		current = next

		if next != nil {
			l.Head = current
		}
	}
}

// h1 > h2
func ZipperLinkedLists(h1, h2 *Node) *Node {
	if h1 == nil {
		return nil
	}

	counter := 1

	c1 := h1
	c2 := h2

	l := LinkedList{}

	for c1 != nil || c2 != nil {

		if counter%2 != 0 {
			l.Insert(c1.Value)
			c1 = c1.Next
			counter += 1
			continue
		}

		l.Insert(c2.Value)
		c2 = c2.Next
		counter += 1

	}

	return l.Head
}

func main() {
	list := LinkedList{}
	list.Insert(12)
	list.Insert(22)
	list.Insert(32)
	list.Insert(54)
	list.Insert(90)
	list.Insert(331)
	list.Insert(2121)
	list.Insert(441)

	list.Delete(12)

	current := list.Head
	sum := 0

	for current != nil {
		sum += current.Value
		current = current.Next
	}

	fmt.Println(sum)

	fmt.Println(list.Get(0))
	fmt.Println(list.Get(3))

	list2 := LinkedList{}
	list2.Insert(0)
	list2.Insert(1)
	list2.Insert(2)
	list2.Insert(3)
	list2.Reverse()
	current = list2.Head

	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}

	l1 := LinkedList{}
	l1.Insert(2)
	l1.Insert(4)
	l1.Insert(6)
	l1.Insert(8)

	l2 := LinkedList{}
	l2.Insert(5)
	l2.Insert(15)
	l2.Insert(25)
	l2.Insert(35)

	zip := ZipperLinkedLists(l1.Head, l2.Head)

	fmt.Println("zipper node")
	for zip != nil {
		fmt.Println(zip.Value)
		zip = zip.Next
	}
}
