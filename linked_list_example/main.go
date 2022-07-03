package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func NewSinglyLinkedList(nums []int) SinglyLinkedList {
	dummyNode := &Node{}
	currNode := dummyNode
	for _, n := range nums {
		currNode.next = &Node{n, nil}
		currNode = currNode.next
	}

	return SinglyLinkedList{dummyNode.next}
}

func main() {
	linkedList := NewSinglyLinkedList([]int{5})
	fmt.Println(linkedList)
}
