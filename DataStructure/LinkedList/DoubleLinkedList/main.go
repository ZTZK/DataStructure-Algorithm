package main

import (
	"fmt"
)

type DoubleLinkedListNode struct {
	content  int                   // Store the data
	preIndex *DoubleLinkedListNode // Point to the previous one node
	proIndex *DoubleLinkedListNode // Point to the next one node
}

type DoubleLinkedList struct {
	head *DoubleLinkedListNode // Head node
	tail *DoubleLinkedListNode // Tail node
	size int                   // Size of the list
}

func (list *DoubleLinkedList) initialize() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *DoubleLinkedList) get(index int) *DoubleLinkedListNode {
	if list.size <= index || index < 0 { // If the index is larger or equal to the size of the list or illegal, then return nil
		return nil
	}

	temp := list.head
	for i := 0; i != index; i++ {
		temp = temp.proIndex
	}

	return temp
}

func (list *DoubleLinkedList) append(newNode *DoubleLinkedListNode) {
	if newNode == nil {
		return
	} // If the new node is nil, do nothing
	if list.head == nil { // If the new node will be the first node of the list
		list.head = newNode
		list.tail = newNode
		list.size++
		newNode.preIndex = nil
		newNode.proIndex = nil
	} else { // If the new node will not be the first node of the list
		newNode.preIndex = list.tail
		newNode.proIndex = nil
		list.size++
		list.tail.proIndex = newNode
		list.tail = newNode
	}
}

func (list *DoubleLinkedList) insert(index int, newNode *DoubleLinkedListNode) {
	if newNode == nil { // The the inserted is nil, then do nothing
		return
	}
	if list.head == nil { // If the inserted node is the first node of the list
		list.append(newNode)
		return
	}
	if index >= list.size { // If the index is larger or equal to the size of the list, then append it to the tail
		list.append(newNode)
		return
	}

	// Below is inserting a node to the inner part of the list
	position := list.get(index)
	newNode.proIndex = position
	newNode.preIndex = position.preIndex
	position.preIndex.proIndex = newNode
	position.preIndex = newNode
	list.size++

}

func (list *DoubleLinkedList) delete(index int) {
	if index < 0 || index >= list.size {
		return
	}
	if index == 0 { // If deleting the head node of an non-empty list
		list.head = list.head.proIndex
		list.head.preIndex = nil
		list.size--
		return
	}
	if index == list.size-1 { // If deleting the last node of an non-empyt list
		list.tail = list.tail.preIndex
		list.tail.proIndex = nil
		list.size--
		return
	}
	temp := list.get(index)
	temp.preIndex.proIndex = temp.proIndex
	temp.proIndex.preIndex = temp.preIndex
	list.size--

}

func (list *DoubleLinkedList) display() {
	if list.head == nil {
		return
	}
	temp := list.head
	temp = list.head
	for temp != nil {
		fmt.Println(temp.content)
		temp = temp.proIndex
	}
}

func newNode(data int) *DoubleLinkedListNode {
	node := new(DoubleLinkedListNode)
	node.content = data
	node.preIndex = nil
	node.proIndex = nil
	return node
}

func main() {
	// This file shows the basic operations of double linked list

	// Firstly, we initialize a list
	newList := new(DoubleLinkedList)

	// Next, we create a new node
	node1 := newNode(3)

	// Then, we append the node1 to the newList
	newList.append(node1)

	// Next, we display the list
	//newList.display()

	// Next, we append node2
	node2 := newNode(1)
	newList.append(node2)
	node3 := newNode(9)
	newList.append(node3)
	newList.insert(1, newNode(7))
	newList.delete(2)
	newList.display()

}
