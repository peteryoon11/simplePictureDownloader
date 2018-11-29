package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readLine("./defaultSource/test.txt")
}
func readLine(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		temp := scanner.Text()
		//fmt.Println(scanner.Text())
		//	fmt.Println(temp)
		parsingImgPath(temp)
	}
}
func parsingImgPath(temp string) {
	if strings.Contains(temp, "<img") {
		fmt.Println(temp)
	}

}
