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

func main() {
	myList := &Node{-100, nil}
	myList.append(100)
	myList.append(101)
	myList.append(102)
	myList.append(103)

	fmt.Println("============= Forward ============")
	myList.traverse()
	fmt.Println("============= REVERSE ============")
	myList.reverse()
	fmt.Println("============= End ============")

}
