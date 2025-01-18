package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Add(x, y float64) {
	p.x, p.y = p.x+x, p.y+y
}

func (p Point) OffsetOf(p1 Point) (x float64, y float64) {
	x, y = p.x-p1.x, p.y-p1.y
	return
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func receivers() {
	q := Point{3, 4}
	p := Point{1, 1}

	fmt.Println(p.Distance(q))

	distanceFromP := p.Distance // it copies the pointer to p, not the value of p
	fmt.Printf("%T\n", distanceFromP)
	//fmt.Printf("%T\n", Point.Distance)
	fmt.Println(distanceFromP(q))

	p = Point{2, 2}
	fmt.Println(distanceFromP(q)) // No change if (p Point)

}
