package main

import (
	"fmt"
	"path/filepath"
)

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hasf of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	Length int
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hasf of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type Filenamer interface {
	Filename() string
}

func pairs() {
	p := Pair{"/usr", "0xfdfe"}

	pl := PairWithLength{Pair{"usr/lib/", "0xdead"}, 133}

	var fn Filenamer = PairWithLength{Pair{"usr/lib/", "0xdead"}, 133}

	fmt.Println(p)
	fmt.Println(pl)

	fmt.Println(fn.Filename())
}
