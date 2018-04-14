package main

import (
	"sort"
	"fmt"
)

func main() {
	a:=[]int{2,3,6,7,1,-1,100,35}
	sort.Ints(a)
	for _, v:=range a{
		fmt.Println(v)
	}
}
