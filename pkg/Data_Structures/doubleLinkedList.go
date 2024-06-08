// Copyright (c) avijit bhattacharjee 2024

package Data_Structures

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func insertStart(root *Node, val int) *Node {
	var newNode = &Node{
		Val:   val,
		Right: root,
		Left:  nil,
	}

	if root != nil {
		root.Left = newNode
	}

	return newNode
}

func insertLast(root *Node, val int) *Node {
	var newNode = &Node{
		Val:   val,
		Right: nil,
		Left:  nil,
	}

	if root == nil {
		return newNode
	}

	var dummyRoot = root

	for dummyRoot.Right != nil {
		dummyRoot = dummyRoot.Right
	}

	dummyRoot.Right = newNode
	newNode.Left = dummyRoot

	return root
}

func deleteStart(root *Node) *Node {
	if root == nil {
		return nil
	}

	fmt.Println(root.Val, " is being deleted")
	root = root.Right
	if root != nil {
		root.Left = nil
	}

	return root
}

func deleteLast(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Right == nil {
		fmt.Println(root.Val, " is being deleted")
		return nil
	}

	current := root
	for current.Right.Right != nil {
		current = current.Right
	}
	fmt.Println(current.Right.Val, " is being deleted")
	current.Right = nil

	return root
}

func displayDLL(root *Node) {
	if root == nil {
		return
	}

	fmt.Println("Displaying DLL")
	for root != nil {
		fmt.Print(root.Val, " ")
		root = root.Right
	}
	fmt.Println("Displayed DLL")
}

func ImplementDoubleLinkedList() {
	var root *Node

	root = insertStart(root, 3)
	root = insertStart(root, 2)
	root = insertStart(root, 1)
	root = insertLast(root, 4)
	root = insertLast(root, 5)

	displayDLL(root)

	root = deleteLast(root)
	root = deleteStart(root)

	displayDLL(root)
}
