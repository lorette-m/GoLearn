package main

import (
	"fmt"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func testFib() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}

	d, g := fib(), fib()

	fmt.Println(d(), d(), d(), d(), d())
	fmt.Println(g(), g(), g(), g(), g())
}

func do(d func()) {
	d()
}

func checkLoopIncAddress1() {
	var i int
	for i = 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

		do(v)
	}
}

func checkLoopIncAddress2() {
	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

		do(v)
	}
}

func last() { //HAHAHAHAAHAH IT IS AN ERROR ISNT IT? YOU R FUCKING FAGOT IT'S COMPILER FUCKIN ERROR LOOOOOOOOOOOL DUDE
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i // onestly fix doesn't fix this closure issue
		s[i] = func {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}
	// needed to clarify how to work with this mechanic and ensuing issues on Go 1.21 and higher

	for i := 0; i < 4; i++ {
		s[i]()
	}
}

func main() {
	//testFib()
	//checkLoopIncAddress1()
	/*Output:
	  0 @ 0xc00000a0c8
	  1 @ 0xc00000a0c8
	  2 @ 0xc00000a0c8
	  3 @ 0xc00000a0c8
	*/
	//checkLoopIncAddress2()
	/*Output:
		    0 @ 0xc00000a0c8
	        1 @ 0xc00000a0e8
	        2 @ 0xc00000a110
	        3 @ 0xc00000a118
	*/

}
