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

}
