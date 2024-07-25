package doublyLinkedList

type Node struct {
	Value interface{}
}

type DoublyLinkedList struct {
	currentIndex int
	container    []Node
}

func (l *DoublyLinkedList) Len() int {
	return len(l.container)
}

func (l *DoublyLinkedList) First() interface{} {
	return l.container[0].Value
}

func (l *DoublyLinkedList) Last() interface{} {
	return l.container[l.Len()-1].Value
}

func (l *DoublyLinkedList) PushFront(value interface{}) {
	n := Node{Value: value}
	l.container = append([]Node{n}, l.container...)
}

func (l *DoublyLinkedList) PushBack(value interface{}) {
	n := Node{Value: value}
	l.container = append(l.container, n)
}

func (l *DoublyLinkedList) Remove(index int) {
	if index < 0 || index > l.Len()-1 {
		return
	}
	l.container = append(l.container[:index], l.container[index+1:]...)
}

func (l *DoublyLinkedList) Item() (Node, string) {
	if l.Len() > 0 {
		return l.container[l.currentIndex], ""
	}
	return Node{}, "empty list"
}

func (l *DoublyLinkedList) Value() (interface{}, string) {
	if l.Len() > 0 {
		return l.container[l.currentIndex].Value, ""
	}
	return nil, "empty list"
}

func (l *DoublyLinkedList) Next() (Node, string) {
	l.currentIndex++
	if l.currentIndex > l.Len()-1 {
		l.currentIndex = 0
	}
	return l.Item()
}

func (l *DoublyLinkedList) Prev() (Node, string) {
	l.currentIndex--
	if l.currentIndex < 0 {
		l.currentIndex = l.Len() - 1
	}
	return l.Item()
}
