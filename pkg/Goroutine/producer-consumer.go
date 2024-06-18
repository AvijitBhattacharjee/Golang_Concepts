// Copyright (c) avijit bhattacharjee 2024

package Goroutine


import (
	"fmt"
	"sync"
	"time"
)

func SyncProducerConsumer() {

	fmt.Println("Implememting sync")

	var num = make(chan int)
	
	var wg sync.WaitGroup

	wg.Add(2)
	go producerSync(num, &wg)
	go consumerSync(num, &wg)

	wg.Wait()
	close(num)
}

func producerSync(num chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i:=0;i<5;i++ {
		num <- i
		fmt.Println("Produced = ", i)
		time.Sleep(time.Second)
	}
}

func consumerSync(num chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i:=0;i<5;i++ {
		val := <-num
		fmt.Println("Produced = ", val)
		time.Sleep(time.Second)
	}

}