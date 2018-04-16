package main

import "fmt"

type Data struct {


}

func (d Data) String() string{
	return "data"
}

func main() {
	d:=Data{}
	fmt.Println(d)
}
