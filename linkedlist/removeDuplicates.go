package linkedlist

import "fmt"

func DeleteDuplicates(head *Node) *Node {
	current := head
	for current != nil && current.next != nil {
		if current.data == current.next.data {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
	return head
}

func PrintList(head *Node) {
	current := head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

}
