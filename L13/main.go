package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func B() string {
	_, file, line, _ := runtime.Caller(1)

	idx := strings.LastIndexByte(file, '/')

	return "=> " + file[idx+1:] + ":" + strconv.Itoa(line)
}

func A() string {
	return B()
}

func regexFind() {
	te := "aba abba abbba"
	re := regexp.MustCompile("b+")
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)

	fmt.Println(mm)
	fmt.Println(id)

	for _, d := range id {
		fmt.Println(te[d[0]:d[1]])
	}
}

func regexReplace() {
	te := "abba abba abbba"
	re := regexp.MustCompile("b+")
	up := re.ReplaceAllStringFunc(te, strings.ToUpper)

	fmt.Println(up)
}

func main() {
	/*test := "Here is $1 which is $2!"

	test = strings.ReplaceAll(test, "$1", "honey")
	test = strings.ReplaceAll(test, "$2", "tasty")

	fmt.Println(test)*/

	//fmt.Println(A())

	//regexFind()

	//regexReplace() // aBBa aBBa aBBBa

	//phoneMask1()

	phoneMask2()
}
