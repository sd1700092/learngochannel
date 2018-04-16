package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string `xml:"name,attr"`
	Age  int `xml:"age"`
}

func main() {
	p := person{Name: "davy", Age: 18}
	var data []byte
	var err error
	if data, err = xml.Marshal(p); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(data))
	}

	if data, err = xml.MarshalIndent(p, "", "	"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	p2 := new(person)
	fmt.Println(p2)
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}
