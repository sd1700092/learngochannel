package main

import (
	"practice/learnchannel/pipeline"
	"fmt"
)

func main() {
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))
	/*	for {
			if num, ok := <-p; ok {
				fmt.Println(num)
			} else {
				break
			}
		}*/
	for v := range (p) { //如果p里面没有数据，程序会自然而然地等待，不需要我们干预
		fmt.Println(v)
	}
}
