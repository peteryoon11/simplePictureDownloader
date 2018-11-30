package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//readLine("./defaultSource/test.txt")

	path := "./defaultSource/testparsing.txt"
	//	path := "c:/dev/private/golang/simple-tool/test_Code/webSourceParsing/defaultSource/testparsing.txt"
	readLine(path)
	/* fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", "")) */

	//readLinever2(path)
	//readLine("./defaultSource/testparsing.txt")
}
func readLine(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		temp := scanner.Text()
		//fmt.Println(scanner.Text())
		fmt.Println("no = ", i)
		i++
		//fmt.Println(temp)
		parsingImgPath(temp)
	}
}
func parsingImgPath(temp string) {
	if strings.Contains(temp, "abc") {
		fmt.Println(temp)
	}

}
func readLinever2(path string) {
	//file, err := os.Open("/path/to/file.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
