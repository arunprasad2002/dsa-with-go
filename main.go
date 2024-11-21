package main

import "example.com/dsa/linkedlist"

func main() {
	ll := &linkedlist.Linkedlist{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(40)

	ll.Display()
}
