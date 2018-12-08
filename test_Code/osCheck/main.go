package main

import (
	"fmt"
	"os"
	"path"
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
func TestFunc(filepath string) {
	filepathOnlyPath, _ := path.Split(filepath)
	if _, err := os.Stat(filepathOnlyPath); !os.IsNotExist(err) {
		//!os.IsNotExist(err) for window
		// os.IsNotExist(err)  for mac
		// path/to/whatever does not exist
		//fmt.Println("filepathOnlyPath = ", filepathOnlyPath)
		// 이렇게 여러번 확인 할 필요가 있나.. 싶은데.. 나중에 다시 체크 하자.
		err := os.Mkdir(filepathOnlyPath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("filepath is exist")
	}
}
