package main

/*
     Fan-In Pattern: Multiple workers send messages to a single channel, and the main routine processes them.
     Visualization of the Fan-In pattern:
    
           +-------+       +-------------+
           | Worker|       |             |
           |   1   | ----> |             |
           +-------+       |             |
                           |             |
           +-------+       |             |
           | Worker|       |             |
           |   2   | ----> |    Channel  |
           +-------+       |     (ch)    | ----> Main Go Routine
                           |             |       Receives Messages
           +-------+       |             |
           | Worker|       |             |
           |   3   | ----> |             |
           +-------+       |             |
                           |             |
           +-------+       |             |
           | Worker|       |             |
           |   4   | ----> |             |
           +-------+       |             |
                           |             |
           +-------+       +-------------+
           | Worker|
           |   5   | ---->
           +-------+
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, exact bool, ch chan string) {
	duration := time.Duration(10+rand.Intn(31)) * time.Second

	if exact {
		duration = 5 * time.Second
	}

	time.Sleep(duration)

	message := fmt.Sprintf("Data from worker **%d** after %v", id, duration)
	ch <- message
}



func main() {
	rand.NewSource(time.Now().UnixNano())
	ch := make(chan string, 1) // can hold up to 1 message
	
	go worker(1, false, ch)
	go worker(2, true, ch)
	go worker(3, false, ch)
	go worker(4, true, ch)
	go worker(5, false, ch)
	go worker(6, false, ch)
	go worker(7, true, ch)

	for i := 1; i <= 7; i++ {
		msg := <-ch
		fmt.Printf("Main received: %s\n", msg)
	}
}
