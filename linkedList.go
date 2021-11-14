/*
	forvaidya@gmail.com
	+91-9740500144
*/
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (currNode *Node) traverse() {
	if currNode == nil {
		return
	}
	fmt.Println(currNode.data)
	if currNode.next != nil {
		currNode.next.traverse()
	}
}

func (currNode *Node) reverse() {
	if currNode.next != nil {
		currNode.next.reverse()
	}
	fmt.Println(currNode.data)
}

func (currNode *Node) append(data int) {
	if currNode.next == nil {
		currNode.next = &Node{data, nil}
	} else {
		currNode.next.append(data)
	}
}

func (currNode *Node) reverse_new() *Node {
	var tmpNode *Node
	prevNode := currNode

	if currNode == nil {
		return tmpNode // return nil
	}

	if currNode.next == nil {
		return currNode // only one node ::return node itself
	}

	tmpNode = currNode.next.reverse_new()
	tmpNodex := tmpNode

	/*
		why this snippet
		because I have to attach temp node to extreme end.
		This is a hog

	*/
	for tmpNodex.next != nil {
		tmpNodex = tmpNodex.next
	}

	/* end why this snippet */

	tmpNodex.next = prevNode
	tmpNodex.next.next = nil
	return tmpNode

}

func (currNode *Node) reverse_iter() *Node {

	var next *Node
	var prev *Node
	var current *Node

	current = currNode

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev

}

func main() {
	myList := &Node{-9, nil}
	myList.append(1)
	myList.append(2)
	myList.append(3)
	myList.append(4)

	myList.reverse_iter().traverse()

	/*

		fmt.Println("============= Forward ============")
		myList.traverse()

		fmt.Println("============= Reverse ============")
		myList.reverse_new().traverse()

		fmt.Println("============= END ============")

	*/
}
