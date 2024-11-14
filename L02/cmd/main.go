package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}
