package main

import (
	"fmt"
	"os"
)

func main() {
	filepathOnlyPath := "./temp"
	if _, err := os.Stat(filepathOnlyPath); !os.IsNotExist(err) {
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
