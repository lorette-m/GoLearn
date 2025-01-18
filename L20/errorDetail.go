package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

func XYZ(a int) error { // error instead of *errFoo
	return nil
}

func errorDetail() {
	// err := XYZ(1) // err would be *errFoo
	var err error = XYZ(1)

	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok!")
	}
}
