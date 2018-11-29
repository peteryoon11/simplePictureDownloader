package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(scanner.Text())
	}
}
