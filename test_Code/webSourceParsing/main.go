package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	//readLine("./defaultSource/test.txt")

	//	path := "./defaultSource/testparsing.txt"
	path := "./defaultSource/test.txt"
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
	//i := 0
	for scanner.Scan() {
		temp := scanner.Text()
		//fmt.Println(scanner.Text())
		//fmt.Println("no = ", i)
		//i++
		//fmt.Println(temp)
		parsingImgPath(temp)
	}
}
func parsingImgPath(temp string) {
	if strings.Contains(temp, "<img") {
		//fmt.Println(temp)
		parsingImgSrc(temp)
	}

}
func parsingImgSrc(temp string) {
	if strings.Contains(temp, "src") {
		fmt.Println(temp)
		//parsingImgSrc(temp)
	}
}
func FindToken() {
	
	z := html.NewTokenizer(response.Body)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				fmt.Println("We found a link!")
			}
		}
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
