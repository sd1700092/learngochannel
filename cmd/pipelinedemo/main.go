package main

import (
	"fmt"
	"learngochannel/pipeline"
	"os"
	"bufio"
)

func main() {
	const filename = "large.in"
	//fmt.Println("rand int = ", rand.Int())
	//const filename = "small.in"
	//const n = 64
	const n = 100000000
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	//for v := range (p){
	//	fmt.Println(v)
	//}
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	//mergeDemo()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println("value read = ", v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeDemo() {
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

