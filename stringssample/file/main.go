package main

import (
	"io/ioutil"
	"encoding/xml"
	"bytes"
	"fmt"
)

func main() {
	content, err := ioutil.ReadFile("vsproj.csproj")
	decoder := xml.NewDecoder(bytes.NewBuffer(content))
	var t xml.Token
	var inItemGroup bool
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			//fmt.Println(name)
			if inItemGroup {
				if name == "Compile" {
					fmt.Println(getAttributeValue(token.Attr, "Include"))
				}
			} else {
				if name == "ItemGroup" {
					inItemGroup = true
				}
			}
		case xml.EndElement:
			if inItemGroup{
				if token.Name.Local=="ItemGroup" {
					inItemGroup = false
				}
			}
		}
	}
}

func getAttributeValue(attr []xml.Attr, name string) string {
	for _, a:= range attr{
		if a.Name.Local == name{
			return a.Value
		}
	}
	return ""
}
