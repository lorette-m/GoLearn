package main

import (
	"fmt"
	"io"
	"os"
)

// UNIX cat

func UnixCatFunc() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue // kill or continue?
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		file.Close()
	}
}
