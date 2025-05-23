package main

import "example.com/dsa/linkedlist"

func main() {

	linkedList := linkedlist.LinkedList{}
	linkedList.Add(10)
	linkedList.Add(20)
	linkedList.Add(100)

	linkedList.Remove(100)

	linkedList.Display()
}
