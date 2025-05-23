package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Add(data int) {
	newNode := &Node{
		data: data,
	}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func (ll *LinkedList) Display() {
	currentNode := ll.head

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}

}

func (ll *LinkedList) Remove(data int) {
	if ll.head == nil {
		return // List is empty
	}

	// Special case: the head needs to be removed
	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}

	currentNode := ll.head

	for currentNode.next != nil {
		if currentNode.next.data == data {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}
