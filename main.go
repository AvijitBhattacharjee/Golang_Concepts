// Copyright (c) avijit bhattacharjee 2024

package main

import (
	"github.com/avijit/pkg/Data_Structures"
	"github.com/avijit/pkg/JSON_Handling" 
	"github.com/avijit/pkg/Rest_API" 
)

func main() {
	Data_Structures.PrintArray()
	JSON_Handling.Marshall()
	JSON_Handling.Unmarshall()
	Rest_API.Start()
}
