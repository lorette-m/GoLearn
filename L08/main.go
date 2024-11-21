package main

import "fmt"

func do(m1 *map[int]int) {
	//b[0] = 0
	//fmt.Printf("b@ %p\n", b)
	//return b[1]
	(*m1)[3] = 0
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("m1", *m1)
}

func main() {
	m := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("m", m)
	do(&m)
	//a := []int{1, 2, 3}
	//fmt.Printf("a@ %p\n", a)
	//v := do(a)

	fmt.Println("m", m)
}
