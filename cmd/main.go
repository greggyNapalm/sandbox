package main

import (
	"fmt"
	"runtime"
)

var version string = "0.0.2"

func main() {
	fmt.Println("hello platform:", runtime.GOOS, runtime.GOARCH)
	fmt.Println("version:v" + version)
}
