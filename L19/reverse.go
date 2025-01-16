package main

import (
	"fmt"
	"sort"
)

func reverse() {
	entries := []string{"charlie", "able", "dog", "baker"}

	sort.Sort(sort.Reverse(sort.StringSlice(entries)))

	fmt.Println(entries)
}
