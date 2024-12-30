package linkedlist

// MergeTwoLists merges two sorted linked lists into a new sorted linked list
func MergeTwoLists(list1 *Node, list2 *Node) *Node {
	newList := &Linkedlist{}

	list1_current := list1
	list2_current := list2

	// Merge the two lists
	for list1_current != nil && list2_current != nil {
		if list1_current.data <= list2_current.data {
			newList.Add(list1_current.data)
			list1_current = list1_current.next
		} else {
			newList.Add(list2_current.data)
			list2_current = list2_current.next
		}
	}

	// Add remaining nodes from list1, if any
	for list1_current != nil {
		newList.Add(list1_current.data)
		list1_current = list1_current.next
	}

	// Add remaining nodes from list2, if any
	for list2_current != nil {
		newList.Add(list2_current.data)
		list2_current = list2_current.next
	}

	return newList.Head
}
