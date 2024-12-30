package main

import (
	"example.com/dsa/linkedlist"
)

func main() {
	list := linkedlist.Linkedlist{}
	list.Add(10)
	list.Add(10)
	list.Add(20)
	list.Add(30)
	newHead := linkedlist.DeleteDuplicates(list.Head)

	linkedlist.PrintList(newHead)

}
