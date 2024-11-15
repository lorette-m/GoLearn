package main

/*
var a int // явное определение (explicitly)

var ( // неявное определение (implicitly)
	b = 2    // integer
	f = 2.01 // float64
)

c := 2 // short declaration operator
*/
import (
	"fmt"
)

func typeTh() {
	a := 2
	b := 2.1

	fmt.Printf("a: %T %v\n", a, a)
	fmt.Printf("b: %T %v\n", b, b)
	/*
		Output:
		a: int 2
		b: float64 2.1
	*/
	// Syntax optimization
	fmt.Printf("a: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)
	/*
		Output:
		a:      int 2
		b:  float64 2.1
	*/
	a = int(b)
	fmt.Printf("a: %8T %[1]v\n", a)
	b = float64(a)
	fmt.Printf("b: %8T %[1]v\n", b)

	/*const (
		a = 1        // int
		b = 2 * 1024 // 2048
		c = b << 3   // 16384

		g uint8 = 0x07     // 7
		h uint8 = g & 0x03 // 3

		s = "a string"
		t = len(s) // 8
		u = s[2:]  // SYNTAX ERROR
	)*/
}
