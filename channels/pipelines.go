package main

import (
	"fmt"
	"time"
)

func stage1(out chan<- string, duration time.Duration) {
	for i := 1; i <= 5; i++ {
		time.Sleep(duration)
		fmt.Println("(ch1 <- s1) procesing message : ", i)
		out <- fmt.Sprintf("(ch1 <- s1) sending message #%d", i)
	}

	close(out)
}

func stage2(in <-chan string, out chan<- string, duration time.Duration) {
	for m := range in {
		time.Sleep(duration)
		out <- fmt.Sprintf("(ch2 <- s2) sending message : %s", m)
	}

	close(out)
}

func stage3(in <-chan string) {
	for num := range in {
		fmt.Printf("(s3 <-ch2): %s\n", num)
	}

}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	fmt.Println("starting the stages")

	go stage1(ch1, 1*time.Second)
	go stage2(ch1, ch2, 1*time.Second)

	stage3(ch2)

}
