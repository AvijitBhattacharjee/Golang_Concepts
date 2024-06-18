// Copyright (c) avijit bhattacharjee 2024

package Goroutine

import (
	"fmt"
)

func ArraySum() {

	var arr = []int{1,2,3,4,5,6,7,8,9,10}
	var firstPartSum = make(chan int, len(arr)/2)
	var secondPartSum = make(chan int, len(arr)/2)


	go firstSum(arr[:len(arr)/2], firstPartSum)
	go secondSum(arr[len(arr)/2:], secondPartSum)

	fmt.Println("This is the first part sum = ", <-firstPartSum)
	fmt.Println("This is the second part sum = ", <-secondPartSum)
}

func firstSum(arr []int, result chan int) {

	sum := 0

	for _, value := range arr {
		sum+= value
	}
	result <- sum
	close(result)

}

func secondSum(arr []int, result chan int) {

	sum := 0

	for _, value := range arr {
		sum+= value
	}
	result <- sum
	close(result)

}