package main

import "fmt"

func testSlices() {
	var s []int
	t := []int{}
	u := make([]int, 5)
	v := make([]int, 0, 5)

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	// 0, 0, []int,  true, []int(nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	// 0, 0, []int, false, []int{}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	// 5, 5, []int, false, []int{0, 0, 0, 0, 0}
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)
	// 0, 5, []int, false, []int{}
}

func main() {
	//testSlices()
	a := [...]int{1, 2, 3}
	b := a[:1]

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := a[0:2:2] //WTF c =  [1 2] :3
	fmt.Println("c = ", c)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println(len(c))
	fmt.Println(cap(c))

	d := a[0:1:1] // [i:j:k] len j-i cap k-i

	fmt.Println("d = ", d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	//e := d[0:2]
	//fmt.Println("e = ", e) // error

	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
	// the same first element address

	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
	// through slice it's possible to mutate an array...
	// if add third colon c := a[0:2:2] address will change after appending because of reallocation

	c[0] = 9
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
	// but if place it before memory reallocation, array a will change with slice c

}
