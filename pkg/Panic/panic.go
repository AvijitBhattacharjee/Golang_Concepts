// Copyright (c) avijit bhattacharjee 2024

package Panic


import (
	"fmt"
)

func Panic() {
	defer fmt.Println("defer in function 1")
	func2()
}

func func2() {
	defer fmt.Println("defer In function 2")
	panic("panic attack")
	fmt.Print("After panic")
}
