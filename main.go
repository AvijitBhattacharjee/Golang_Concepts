// Copyright (c) avijit bhattacharjee 2024

package main

import (
	"github.com/avijit/pkg/Data_Structures"
	"github.com/avijit/pkg/JSON_Handling" 
	// "github.com/avijit/pkg/Rest_API" 
	"github.com/avijit/pkg/Goroutine"
	"github.com/avijit/pkg/Algorithm/searching"
	"github.com/avijit/pkg/Panic"
	"github.com/avijit/pkg/Variadic"
	"github.com/avijit/pkg/Data_Type"
)

func main() {
	Data_Structures.PrintArray()
	JSON_Handling.Marshall()
	JSON_Handling.Unmarshall()
	//Rest_API.Start()
	Data_Structures.ImplementStack()
	Data_Structures.ImplementQueue()
	Data_Structures.ImplementDoubleLinkedList()
	Data_Structures.ImplementSingleLinkedList()

	Goroutine.ImplementChannel() 
	Goroutine.ImplementInterface()
	Goroutine.ImplementConcurrency()

	Goroutine.ArraySum()
	Goroutine.ProducerConsumer()
	Goroutine.SyncProducerConsumer()

	searching.LinearSearch()
	searching.BinarySearch()

	Panic.Panic()
	Panic.Recover()
	Variadic.Variadic()
	Data_Type.Rune()


}
