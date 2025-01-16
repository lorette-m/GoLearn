package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*type Stringer interface {
	String() string
}*/

type IntSlice []int

func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func m1() {
	var v IntSlice = []int{1, 2, 3}
	var s fmt.Stringer = v

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", s)
	fmt.Printf("%T %[1]v\n", v)
}
