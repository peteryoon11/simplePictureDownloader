package main

import (
	"fmt"
	"path"
)

func main() {
	temp := "./temp/temp1/01.jpg"
	fmt.Println(temp)
	//fmt.Println(strings.Split(temp, "/")[0])
	/* 	for _, item := range strings.Split(temp, "/") {
		fmt.Println(item)
	} */
	fmt.Println(path.Split(temp))
}
