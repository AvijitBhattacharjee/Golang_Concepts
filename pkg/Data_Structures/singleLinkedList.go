// Copyright (c) avijit bhattacharjee 2024

package Data_Structures

import (
	"fmt"
)

type SingleNode struct {
	Val  int
	Next *SingleNode
}

func insertStartSLL(root *SingleNode, val int) *SingleNode {
	var newSingleNode = &SingleNode{
		Val:  val,
		Next: root,
	}

	fmt.Println(newSingleNode.Val, " has been added")
	return newSingleNode
}

func insertEndSLL(root *SingleNode, val int) *SingleNode {
	var newSingleNode = &SingleNode{
		Val:  val,
		Next: nil,
	}

	if root == nil {
		return newSingleNode
	}

	current := root
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newSingleNode

	fmt.Println(newSingleNode.Val, " has been added")
	return root
}

func deleteStartSLL(root *SingleNode) *SingleNode {
	if root == nil {
		return nil
	}

	fmt.Println(root.Val, " has been deleted")
	return root.Next
}

func deleteEndSLL(root *SingleNode) *SingleNode {
	if root == nil {
		return nil
	}

	if root.Next == nil {
		return nil
	}

	current := root
	for current.Next.Next != nil {
		current = current.Next
	}
	fmt.Println(current.Next.Val, " has been deleted")
	current.Next = nil
	return root
}

func displaySLL(root *SingleNode) {
	if root == nil {
		fmt.Println("nothing to show")
		return
	}

	fmt.Println("Displaying Single Linked List")
	for root != nil {
		fmt.Print(root.Val, " ")
		root = root.Next
	}
	fmt.Println("\nDisplayed Single Linked List")
}

func ImplementSingleLinkedList() {
	var root *SingleNode

	root = insertStartSLL(root, 3)
	root = insertStartSLL(root, 2)
	root = insertStartSLL(root, 1)
	root = insertEndSLL(root, 4)
	root = insertEndSLL(root, 5)

	displaySLL(root)

	root = deleteStartSLL(root)
	root = deleteEndSLL(root)

	displaySLL(root)
}
