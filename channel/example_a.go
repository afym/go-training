package main

import (
	"fmt"
	"time"
)

func signalGenerator(label string, size int, signals chan<- int) {
	for i := 1; i <= size; i++ {
		fmt.Println("generating a signal for ", label)
		signals <- i
		time.Sleep(time.Second)
	}

	fmt.Println("closing signal channel for ", label)

	close(signals)
}

func mx(signals <-chan int) {
	for s := range signals {
		fmt.Println("sx -> processing mx singals :  ", s)
		time.Sleep(time.Second * 1)
	}
}

func my(signals <-chan int) {
	for s := range signals {
		fmt.Println("sy -> processing mx singals :  ", s)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	mxSignals := make(chan int, 1)
	mySignals := make(chan int, 1)

	fmt.Println("start of the program")

	go signalGenerator("mx", 12, mxSignals)
	go signalGenerator("my", 9, mySignals)

	go mx(mxSignals)
	go my(mySignals)

	for {
		ox, oy := false, false

		select {
		case mx, ok := <-mxSignals:
			if ok {
				ox = ok
				fmt.Println("processing mx signal #", mx)
			}
		case my, ok := <-mySignals:
			if ok {
				oy = ok
				fmt.Println("processing my signal #", my)
			}
		}

		if !ox && !oy {
			break
		}
	}

	fmt.Println("end of the program")
}
