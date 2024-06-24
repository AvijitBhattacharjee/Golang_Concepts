// Copyright (c) avijit bhattacharjee 2024

package Panic


import (
	"fmt"
)

func Recover() {

	var arr = []int{1,2}
	funcRecover(arr,2)
	fmt.Println("Existing normally")
}

func funcRecover(arr []int, index int) {

	defer handlePanic()
	if index > len(arr)-1 {
		panic("Array out of bound")
	}
	fmt.Println("Exiting from func2")
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic recovered")
	}
}