package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
	"time"
	"fmt"
)

var startTime time.Time
func Init(){
	startTime = time.Now()
}

func ArraySource(a ...int) <-chan int { //这个函数的目的是把int类型的数据都导入到channel里面去
	out := make(chan int)
	//channel是goroutine和goroutine之间的通信，不能自己导，一定要用go func
	go func() {
		for _, v := range (a) {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		// read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done: ", time.Now().Sub(startTime))
		// Sort
		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Now().Sub(startTime))

		// Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime))
	}()
	return out
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan int { // 如果传的是-1，代表读全部文件
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead:=0
		for {
			n, err := reader.Read(buffer) //n返回我读到了多少个字节
			bytesRead+=n
			if n > 0 { //代表已经读到了数据
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize !=-1 && bytesRead >= chunkSize){
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int{
	out:=make(chan int)
	go func(){
		for i:=0;i<count;i++{
			out<-rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m:=len(inputs) / 2
	// merge inputs[0..m) and inputs [m..end)
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}
