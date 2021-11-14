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

func main() {
	myList := &Node{-100, nil}
	myList.append(1)
	myList.append(2)
	myList.append(3)
	myList.append(4)

	myList.append(102)

	myList.append(102)
	myList.append(103)

	fmt.Println("============= Forward ============")
	myList.traverse()

	fmt.Println("============= Reverse ============")
	myList.reverse_new().traverse()

	fmt.Println("============= END ============")

}
