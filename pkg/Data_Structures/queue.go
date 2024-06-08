// Copyright (c) avijit bhattacharjee 2024

package Data_Structures


import ( 
	"fmt"
)

type MyQueue struct {
	myslice []interface{}
}

func pushQueue(myqueue *MyQueue, val int) {

	myqueue.myslice = append(myqueue.myslice, val)
	fmt.Println(val, " has been pushed")
}

func popQueue(myqueue *MyQueue) {
	
	val := myqueue.myslice[0]
	myqueue.myslice = myqueue.myslice[1:]
	fmt.Println(val, " has been popped")
}

func displayQueue(myqueue *MyQueue) {
	
	queueLenght := len(myqueue.myslice)

	for i := 0; i < queueLenght; i++ {
		fmt.Print(myqueue.myslice[i], " ")
	}
	fmt.Println(" ")
}

func ImplementQueue() {

	var myqueue = &MyQueue{}
	pushQueue(myqueue, 11)
	pushQueue(myqueue, 22)
	pushQueue(myqueue, 33)

	displayQueue(myqueue)

	popQueue(myqueue)
	displayQueue(myqueue)
}