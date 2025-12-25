package main

import (
	"fmt"
)

func channelDemo() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine"
	}()

	msg := <-ch

	fmt.Println(msg)
}

func bufferedChannelDemo(){
	
	jobs := make(chan int , 3) //buffer capacity = 3
	jobs <- 1
	jobs <- 2
	jobs <- 3

	close(jobs)

	for j := range jobs {
		fmt.Println("Job:" , j)
	}

}