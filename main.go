// Copyright (c) avijit bhattacharjee 2024

package main

import (
	"github.com/avijit/pkg/Data_Structures"
	"github.com/avijit/pkg/JSON_Handling" 
	"github.com/avijit/pkg/Rest_API" 
	"github.com/avijit/pkg/Goroutine"
)

func main() {
	Data_Structures.PrintArray()
	JSON_Handling.Marshall()
	JSON_Handling.Unmarshall()
	Rest_API.Start()
	Data_Structures.ImplementStack()
	Data_Structures.ImplementQueue()
	Data_Structures.ImplementDoubleLinkedList()
	Data_Structures.ImplementSingleLinkedList()

	Goroutine.ImplementChannel() 
	Goroutine.ImplementInterface()
	Goroutine.ImplementConcurrency()
}
