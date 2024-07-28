package main

/*
Visualization of the Worker Pool pattern:

    +-------------+         +-------+
    |             |         |       |
    |   Task 1    | ---->   |       |
    |             |         |       |
    +-------------+         |       |
                            | Worker|
    +-------------+         |       |
    |             |         |   1   |
    |   Task 2    | ---->   |       |
    |             |         |       |
    +-------------+         |       |
                            +-------+
    +-------------+         +-------+
    |             |         |       |
    |   Task 3    | ---->   | Worker|
    |             |         |   2   |
    +-------------+         |       |
                            +-------+
    +-------------+         +-------+
    |             |         |       |
    |   Task 4    | ---->   | Worker|
    |             |         |   3   |
    +-------------+         |       |
                            +-------+
    +-------------+         +-------+
    |             |         |       |
    |   Task 5    | ---->   | Worker|
    |             |         |   4   |
    +-------------+         |       |
                            +-------+
    
    Task Queue                 Worker Pool
*/


import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, duration time.Duration, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("** calling worker ", id, " -> tasks", len(tasks))
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(duration)
	}
}

func main() {
	tasks := make(chan int, 3) // 
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(1, 4*time.Second, tasks, &wg)
	wg.Add(1)
	go worker(2, 2*time.Second, tasks, &wg)
	wg.Add(1)
	go worker(3, 3*time.Second, tasks, &wg)

	// sending tasks
	tasks <- 10
	tasks <- 11
	tasks <- 29
	tasks <- -7
	tasks <- -90
	tasks <- 30

	close(tasks)

	wg.Wait()
}
