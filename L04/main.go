package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "élite"

	fmt.Printf("%8T % [1]v\n", s)
	fmt.Printf("%8T % [1]v\n", []rune(s))
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))
	/*
		Output:
			string élite
		 []int32 [ 233  108  105  116  101]
		 []uint8 [195 169 108 105 116 101] 6
	*/
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		st := strings.Split(scan.Text(), old)
		t := strings.Join(st, new)

		fmt.Println(t)
	}
}

// Get-Content test.txt | go run . matt sonic
