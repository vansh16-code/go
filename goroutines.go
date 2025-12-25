package main

import (
	"fmt"
	"time"
)

func printNumber(){
	for i :=0 ; i <= 5 ;i++{
		fmt.Println("Goroutine : " , i)
		time.Sleep(200 *time.Millisecond)
	}
}

func goroutineDemo(){
	go printNumber()

	for i :=1 ; i <= 5 ; i++{
		fmt.Println("Main :" ,i)
		time.Sleep(200 *time.Millisecond)
	}
}