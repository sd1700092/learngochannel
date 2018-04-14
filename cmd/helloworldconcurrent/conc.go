package main

import (
	"fmt"
	//"time"
)

func main() {
	c:=make(chan string)
	for i := 0; i < 5000; i++ {
		go printHelloWorld(i, c)
	}
	//time.Sleep(5 * time.Millisecond)
	for {
		msg := <- c
		fmt.Println(msg)
	}
}

func printHelloWorld(i int, c chan string) {
	for {
		c <- fmt.Sprintf("Hello world from go routine %d\n", i)
	}
}
