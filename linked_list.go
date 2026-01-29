package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head           *Node
	linkedListSize int
}

func (l *LinkedList) push(value int) {
	newNode := &Node{
		value: value,
	}

	if l.head == nil {
		l.head = newNode
		l.linkedListSize++
		return
	}

	newNode.next = l.head
	l.head = newNode
	l.linkedListSize++
}

func (l *LinkedList) showLinkedList() {
	var pointer *Node = l.head

	for pointer != nil {
		fmt.Printf("%v -> ", pointer.value)
		pointer = pointer.next
	}

}

func (l *LinkedList) delete(value int) {

}

func (l *LinkedList) size() int {
	return l.linkedListSize
}

func main() {

	linkedList := &LinkedList{}

	linkedList.push(1)
	linkedList.push(2)
	linkedList.push(100)
	linkedList.push(500)

	linkedList.showLinkedList()
}
