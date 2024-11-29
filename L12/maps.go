package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func mapsTest() {
	c := map[string]*Employee{}

	c["Lamine"] = &Employee{"Lamine", 2, nil, time.Now()}

	c["Lamine"].Number++

	c["Matt"] = &Employee{
		Name:   "Andrew",
		Number: 1,
		Boss:   c["Lamine"],
		Hired:  time.Now(),
	}

	/*e.Name = "Andrew"
	e.Number = 1
	e.Hired = time.Now()*/

	//e.Boss = b

	fmt.Printf("%T %+[1]v\n", c["Lamine"])
	fmt.Printf("%T %+[1]v\n", c["Matt"])
}
