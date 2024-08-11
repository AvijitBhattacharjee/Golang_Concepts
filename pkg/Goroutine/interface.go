// Copyright (c) avijit bhattacharjee 2024

package Goroutine

import (
	"fmt"
	"sync"
)

type Animal interface{
	Call(msg string, wg *sync.WaitGroup)
	Eat(food string, wg *sync.WaitGroup)
}

type Cat struct {
	Age int
	Color string
}

type Dog struct {
	Size int
	Weight int
}

func (d *Dog)Call(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("size and call is ", d.Size, msg)
}

func (d *Dog)Eat(food string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("weight and food is ", d.Weight, food)
}

func (c *Cat)Call(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("age and call is ", c.Age, msg)
}

func (c *Cat)Eat(food string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("color and food is ", c.Color, food)
}


func ImplementInterface2() {
	var wg sync.WaitGroup
	
	var animals = []Animal{
		&Cat{Age: 2, Color: "white"},
		&Dog{Size: 3, Weight: 100},
	}
	for _, val := range animals{
		wg.Add(2)
		go val.Call("calling", &wg)
		go val.Eat("milk", &wg)
	}

	wg.Wait()

}
