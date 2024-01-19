package main

import (
	"bytes"
	"fmt"
)

type node struct {
	data       int
	prev, next *node
}

type doublyLinkedList struct {
	head *node
	tail *node
}

func newDoublyLinkedList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (dll *doublyLinkedList) InsertFirst(n *node) {
	if dll.head == nil {
		dll.head = n
		dll.tail = n
		return
	}

	n.next = dll.head
	dll.head.prev = n

	dll.head = n
}

func newNode(data int) *node {
	return &node{
		data: data,
		prev: nil,
		next: nil,
	}
}

func (dll *doublyLinkedList) String() string {
	buf := &bytes.Buffer{}
	defer buf.Reset()

	node := dll.head

	buf.WriteString("nil ")
	buf.WriteString("->")
	for node != nil {
		buf.WriteString(fmt.Sprintf(" %d ", node.data))
		if node.next != nil {
			buf.WriteString("<->")
		}
		node = node.next
	}

	buf.WriteString("-> nil")

	return buf.String()

}

func (dll *doublyLinkedList) AdjustNode(curr *node) {
	if dll.head == curr {
		return
	}

	prev := curr.prev
	next := curr.next

	prev.next = next
	if next != nil {
		next.prev = prev
	}

	curr.prev = nil
	curr.next = dll.head

	dll.head = curr
}

func (dll *doublyLinkedList) RemoveTail() {
	curr := dll.tail
	if curr != nil {
		prev := curr.prev

		if prev != nil {
			prev.next = nil
		}

		dll.tail = prev
	}
	curr = nil
}

func (dll *doublyLinkedList) RemoveNode(curr *node) {
	if dll.head == curr {
		next := curr.next
		next.prev = nil
		dll.head = next
		return
	}

	prev := curr.prev
	next := curr.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	curr = nil
}

/*
func main() {
	dll := newDoublyLinkedList()

	ten := newNode(10)
	tw := newNode(20)
	th := newNode(30)

	dll.InsertFirst(ten)
	dll.InsertFirst(tw)
	dll.InsertFirst(th)

	fmt.Println(dll)

	dll.RemoveNode(ten)
	fmt.Println(dll)
}
*/
