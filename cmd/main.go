package main

import (
	"fmt"
	"runtime"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Println("hello platform:", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("my app %s, commit %s, built at %s", version, commit, date)
}
