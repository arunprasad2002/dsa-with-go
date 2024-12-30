package linkedlist

import "fmt"

// Node represents a single node in the linked list
type Node struct {
	data int
	next *Node
}

// linkedlist represet the linked list

type Linkedlist struct {
	Head *Node
}

// Add appends a new node with the given data to the linked list
func (ll *Linkedlist) Add(data int) {
	newNode := &Node{data: data}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Delete removes the first occurenece of a node with specific value

func (ll *Linkedlist) Delete(data int) {
	if ll.Head == nil {
		fmt.Println("List is empty")
		return
	}
	if ll.Head.data == data {
		ll.Head = ll.Head.next
		return
	}

	current := ll.Head

	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next == nil {
		fmt.Println("Value not found in the list")
		return
	}

	current.next = current.next.next
}

// Display prints all the elements in the linked list

func (ll *Linkedlist) Display() {
	if ll.Head == nil {
		fmt.Println("List is empty")
		return
	}
	current := ll.Head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")

}
