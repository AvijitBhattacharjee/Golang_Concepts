// Copyright (c) avijit bhattacharjee 2024

package Goroutine

import (
	"fmt"
	"sync"
)

func ImplementConcurrency() {

	countRange := 10
	var ch = make(chan bool, 1)
	ch <- true
	var wg sync.WaitGroup

	wg.Add(2)
	go printOdd(ch, &wg, countRange)
	go printEven(ch, &wg, countRange)
	wg.Wait()
}

func printEven(ch chan bool, wg *sync.WaitGroup, countRange int) {

	defer wg.Done()
	for i:=0; i<countRange; i++ {
		if(i%2==0 && <-ch) {
			fmt.Println(i)
			ch <- true
		}
	}
}

func printOdd(ch chan bool, wg *sync.WaitGroup, countRange int) {

	defer wg.Done()
	for i:=0; i<countRange; i++ {
		if(i%2!=0 && <-ch) {
			fmt.Println(i)
			ch <- true
		}
	}
}