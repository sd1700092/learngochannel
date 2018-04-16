package main

import (
	"io"
	"strings"
	"fmt"
	"os"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error){
	p:=make([]byte, num)
	n, err := reader.Read(p)
	if n > 0{
		return p[:n], nil
	}
	return p, err
}

func sampleReadFromString(){
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)
}

func sampleReadStdin()  {
	fmt.Println("please input from stdin:")
	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
}

func sampleReadFile(){
	file, _ := os.Open("main.go")
	defer file.Close()
	data, _ := ReadFrom(file, 90)
	fmt.Println(string(data))
}

func main() {
	//sampleReadFromString()
	//sampleReadStdin()//必须用go run main.go跑，而不是直接启动
	sampleReadFile()
}
