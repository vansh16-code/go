package main

import (
	"fmt"
	"sync"
	"time"
)

func worker (id int , wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("Worker" ,id , "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker", id , "done")
}

func WaitGroupDemo(){
	var wg sync.WaitGroup //creates the waitgroup

	wg.Add(3) // Adding or starting 3 goroutines
	go worker(1,&wg)
	go worker(2,&wg)
	go worker(3,&wg)

	wg.Wait()

	fmt.Println("All workers finished ")
}
