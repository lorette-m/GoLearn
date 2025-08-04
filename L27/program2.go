package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// this is the same program as program1 but with concurrent approach implemented

func collectHashes(pairs <-chan pair, result chan<- results) {
	hashes := make(results)

	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}

	result <- hashes
}

func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool) {
	for path := range paths {
		pairs <- hashFile(path)
	}
	done <- true
}

func searchTree2(dir string, paths chan<- string) error {
	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			paths <- p
		}

		return nil
	}

	return filepath.Walk(dir, visit)
}

func run(dir string) results {
	workers := 2 * runtime.GOMAXPROCS(0)

	paths := make(chan string)
	pairs := make(chan pair)
	done := make(chan bool)
	result := make(chan results)

	for i := 0; i < workers; i++ {
		go processFiles(paths, pairs, done)
	}

	go collectHashes(pairs, result)

	if err := searchTree2(dir, paths); err != nil {
		return nil
	}

	// we must close the paths channel so the workers stop
	close(paths)

	// wait for all the workers to be done
	for i := 0; i < workers; i++ {
		<-done
	}

	// by closing pairs we signal that all the hashes
	// have been collected; we have to do it here AFTER
	// all the workers are done
	close(pairs)

	return <-result
}

func program2() {
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide dir name!")
	}

	if hashes := run(os.Args[1]); hashes != nil {
		for hash, files := range hashes {
			if len(files) > 1 {
				// we will use just 7 chars like git
				fmt.Println(hash[len(hash)-7:], len(files))

				for _, file := range files {
					fmt.Println("  ", file)
				}
			}
		}
	}
}
