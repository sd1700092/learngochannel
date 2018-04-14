package main

import (
	"fmt"
)

func main() {
	chanDemo2()
}

type worker1 struct {
	in   chan int
	done chan bool
}

func chanDemo2() {
	var workers [10]worker1
	for i := 0; i < 10; i++ {
		workers[i] = createWorker1(i)
	}

	for i, worker := range workers {
		//workers[i].in <- i
		worker.in <- 'A' + i
	}

	for i, worker := range workers {
		//workers[i].in <- i
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func createWorker1(id int) worker1 {
	w := worker1{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d, received %d\n", id, n)
		done <- true
	}
}
