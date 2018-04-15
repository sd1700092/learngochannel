package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	s := "hello world"
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))
	fmt.Println(strings.Index(s, "o"))
	ss := "1#2#345"
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)

	fmt.Println(strings.Join(splitedStr, "#"))
	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s,"ld"))

	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.Atoi("711"))

	fmt.Println(strconv.ParseBool("false"))
	//浮点的精度有float32和float64
	fmt.Println(strconv.ParseFloat("3.14", 32))
	fmt.Println(strconv.ParseFloat("3.14", 64))

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 10))
	fmt.Println(strconv.FormatInt(123, 2))
	fmt.Println(strconv.FormatInt(20, 16))
}
