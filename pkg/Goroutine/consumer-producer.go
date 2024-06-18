// Copyright (c) avijit bhattacharjee 2024

package Goroutine

import (
	"fmt"
)

func ProducerConsumer() {

	var limit = 10
	var flag = make(chan bool, 1)
	var num = make(chan int, limit)

	go consumer(num, flag)
	go producer(num, limit)
	<-flag
}

func producer(num chan int, limit int) {

	for i:=0;i<limit;i++ {
		fmt.Println("Produced = ", i)
		num <- i
	}
	close(num)
}

func consumer(num chan int, flag chan bool) {

	for value := range num {
		fmt.Println("Consumed = ", value)
	}
	flag <- true
}
