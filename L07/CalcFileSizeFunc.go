package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// function written in way just reading all data inside
// and in fact this approach is not good sooooooooo func2 below

func CalcFileSize1() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue // kill or continue?
		}

		data, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Println("The file has", len(data), "bytes")

		file.Close()
	}
}

func CalcFileSize2() {
	var tlc, twc, tcc int
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue // kill or continue?
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}

		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()

		tlc += lc
		twc += wc
		tcc += cc
	}
	fmt.Printf("%7d %7d %7d %s\n", tlc, twc, tcc, "Total")
}
