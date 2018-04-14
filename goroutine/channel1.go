package main

import (
	"fmt"
	"time"
)

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func worker(i int, c chan int) {
	//for {
	//	n,ok:=<-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d, received %d\n", i, n)
	//}
	for n := range c {
		fmt.Printf("worker %d, received %d\n", i, n)
	}
}

func createWorker(i int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d, received %c\n", i, <-c)
		}
	}()
	return c
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	c <- 'e'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	time.Sleep(time.Millisecond)
}
