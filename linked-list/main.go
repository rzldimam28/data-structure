package main

import (
	"container/list"
	"fmt"
)

func main() {
	linkedList := list.New()
	linkedList.PushBack("Satu")
	linkedList.PushBack("Dua")
	linkedList.PushBack("Tiga")

	fmt.Println(linkedList)

	fmt.Println("Increasing Looping")
	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("Descreasing Looping")
	for e := linkedList.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	// removing element
	firstNode := linkedList.Front()
	linkedList.Remove(firstNode)

	fmt.Println("Linkedlist after Removing First Node")
	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}