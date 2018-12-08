package main

import (
	"fmt"
	"runtime"
)

func main() {
	//fmt.Println(os.Environ())
	/* 	for _, item := range os.Environ() {
		fmt.Println(item)
	} */
	if runtime.GOOS == "darwin" {
		fmt.Println("Mac OS detected")
	}
	if runtime.GOOS == "linux" { // also can be specified to FreeBSD
		fmt.Println("Unix/Linux type OS detected")
	}
	if runtime.GOOS == "windows" {
		fmt.Println("Windows OS detected")
	}
}
