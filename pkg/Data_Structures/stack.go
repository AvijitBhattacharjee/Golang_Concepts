// Copyright (c) avijit bhattacharjee 2024

package Data_Structures

import (
	"fmt"
)

type MyStack struct {
	myslice []interface{}
}

func push(myStack *MyStack, val int) {
	myStack.myslice = append(myStack.myslice, val)
	fmt.Println(val, " is pushed successfully")
}

func pop(myStack *MyStack) {

	if len(myStack.myslice) == 0 {
		fmt.Println("Stack is already empty")
	}
	val := myStack.myslice[len(myStack.myslice)-1]
	myStack.myslice = myStack.myslice[:len(myStack.myslice)-1]
	fmt.Println(val, " is poped")
}

func display(myStack *MyStack) {

	if len(myStack.myslice) == 0 {
		fmt.Println("Stack is already empty")
	}

	sizeOfStack := len(myStack.myslice)
	for i := sizeOfStack - 1; i >= 0; i-- {
		fmt.Print(myStack.myslice[i], " ")
	}
	fmt.Println("")
}

func ImplementStack() {

	var myStack = &MyStack{}
	push(myStack, 10)
	push(myStack, 20)
	push(myStack, 30)

	display(myStack)

	pop(myStack)

	display(myStack)

}