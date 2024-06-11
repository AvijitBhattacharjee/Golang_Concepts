// Copyright (c) avijit bhattacharjee 2024

package Goroutine


import (
	"fmt"
	"time"
)

func ImplementChannel() {

	fmt.Println("Implementing channel")

	ch := make(chan int, 1)

	go send(ch, 2)
	go receive(ch)
	time.Sleep(time.Second * 1)
	go send(ch, 3)
	go receive(ch)
	time.Sleep(time.Second * 1)
	fmt.Println("this is the capacity = ", cap(ch))
}

func send(ch chan int, val int) {
	ch <- val
	fmt.Println("this is the current length = ", len(ch))
}

func receive(ch chan int) {
	val := <- ch
	fmt.Println("this is the current value = ", val)
}