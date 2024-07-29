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

/*Double-linked list
Purpose: https://en.wikipedia.org/wiki/Doubly_linked_list?Ожидаемые types (pseudocode):?``
// List // container type Len() // list length First() // first Item
Last() // last Item PushFront(v interface{}) // add value to the beginning
PushBack(v interface{}) // add value to the end Remove(i Item) // remove item?Item // list item
Value() interface{} // returns value Nex() *Item // next Item Prev() *Item // previous
Implement a doubly linked list in Go language*/
