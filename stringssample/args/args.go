package main

import (
	"fmt"
	"os"
	"flag"
)

func style(){
	methodPtr := flag.String("method", "default", "method of sample")
	valuePtr:=flag.Int("value", -1, "value of sample")
	flag.Parse()
	fmt.Println(*methodPtr, *valuePtr)
}

func style2(){
	 var method string
	 var value int
	 flag.StringVar(&method, "method", "default", "method of sample")
	 flag.IntVar(&value, "value", -1, "value of sample")
	 flag.Parse()
	 fmt.Println(method, value)
}

func main() {
	strings := os.Args
	for _, str := range strings {
		fmt.Println(str)
	}
	style2()
}
