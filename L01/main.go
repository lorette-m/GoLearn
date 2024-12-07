package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Minute)
}

//  go build -o hello.exe main.go

/*
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o helloENV.exe main.go
*/
