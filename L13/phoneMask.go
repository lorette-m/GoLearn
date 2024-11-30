package main

import (
	"fmt"
	"regexp"
)

var ph = regexp.MustCompile(`\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`)

func phoneMask1() {
	orig := "(214) 514-9548"
	match := ph.FindStringSubmatch(orig)

	fmt.Printf("%q\n", match)

	if len(match) > 3 {
		fmt.Printf("+1 %s-%s-%s\n", match[1], match[2], match[3])
	}
}

const phre = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`

var pfmt = regexp.MustCompile(phre)

func phoneMask2() {
	orig := "call me at (214) 514-9548 today "
	/*
		match := ph.FindStringSubmatch(orig)
		fmt.Printf("%q\n", match)
	*/
	intl := pfmt.ReplaceAllString(orig, "+1 ${1}-${2}-${3}")

	fmt.Println(intl)
}
