package main

import (
	"fmt"
	"l5DoublyLinkedList/doublyLinkedList"
)

func main() {
	list := doublyLinkedList.DoublyLinkedList{}
	item, err := list.Item()
	fmt.Printf("%v %v\n", item, err)
	fmt.Printf("%v %v\n", list, list.Len())
	list.PushFront(2)
	fmt.Printf("%v %v\n", list, list.Len())
	list.PushBack("f")
	fmt.Printf("%v %v %v %v\n", list, list.Len(), list.First(), list.Last())

	list.Remove(2)
	fmt.Printf("%v %v\n", list, list.Len())
	item, err = list.Item()
	fmt.Printf("%v %v\n", item, err)
	val, err := list.Value()
	fmt.Printf("%v %v\n", val, err)
	list.PushBack("q")
	list.PushBack("w")

	item, err = list.Next()
	fmt.Printf("%v %v %v\n", list, item, err)

	item, err = list.Prev()
	fmt.Printf("%v %v %v\n", list, item, err)
	item, err = list.Prev()
	fmt.Printf("%v %v %v\n", list, item, err)
	item, err = list.Next()
	fmt.Printf("%v %v %v\n", list, item, err)

}
