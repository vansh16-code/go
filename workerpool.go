package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)
		time.Sleep(time.Second)
	}
}

func workerPoolDemo() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker2(w, jobs, &wg)
	}

	// Send 6 jobs
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs) // no more jobs

	wg.Wait()

	fmt.Println("All jobs done")
}
