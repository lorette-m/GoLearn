package main

import "fmt"

func FcodesFunc() {
	a, b := 12, 345
	c, d := 1.2, 3.4513

	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%X %X\n", a, b)
	fmt.Printf("%#x %#x\n", a, b)
	fmt.Printf("%f %.2f\n", c, d)

	fmt.Println()

	fmt.Printf("|%9d|%9d|\n", a, b)
	fmt.Printf("|%09d|%09d|\n", a, b)
	fmt.Printf("|%-9d|%-9d|\n", a, b)
	fmt.Printf("|%9f|%9.2f|\n", c, d)

	s := []int{1, 2, 3}
	ar := [3]rune{'a', 'b', 'c'}

	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%T\n", ar)
	fmt.Printf("%q\n", ar)
	fmt.Printf("%v\n", ar)
	fmt.Printf("%#v\n", ar)

	m := map[string]int{"and": 1, "or": 2}

	fmt.Printf("%T\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)

	st := "ar string"
	by := []byte(st)

	fmt.Printf("%T\n", st)
	fmt.Printf("%q\n", st)
	fmt.Printf("%v\n", st)
	fmt.Printf("%#v\n", st)
	fmt.Printf("%v\n", by)
	fmt.Printf("%v\n", string(by))
}
