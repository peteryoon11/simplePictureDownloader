package main

import (
	"fmt"
	"strings"
)

func main() {
	temp := "./temp/01.jpg"
	fmt.Println(temp)
	//fmt.Println(strings.Split(temp, "/")[0])
	for _, item := range strings.Split(temp, "/") {
		fmt.Println(item)
	}
}
