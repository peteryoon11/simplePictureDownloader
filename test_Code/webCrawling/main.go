package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	improvGetWebSourceFunc()
}
func FindToken() {

	//z := html.NewTokenizer(response.Body)

}

func improvGetWebSourceFunc() {

	// Request 객체 생성
	// https://kissme2145.tistory.com/1418?category=634440
	//req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	req, err := http.NewRequest("GET", "https://kissme2145.tistory.com/1418?category=634440", nil)
	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	req.Header.Add("User-Agent", "Crawler")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		//	recover()
	}
	//FindToken(resp.Body)
	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	i := 0
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			fmt.Println(html.ErrorToken)
			return
		case tt == html.StartTagToken:
			t := z.Token()

			//isAnchor := t.Data == "img"
			/* 	if isAnchor {
				//fmt.Println("We found a link!")
				fmt.Println(t.Data)
			} */

			for _, a := range t.Attr {

				//	fmt.Println("key = ", a, " ")
				//if a.Key == "img" {
				if a.Key == "src" {
					//fmt.Println("Found href:", a.Val)
					if strings.Contains(a.Val, "cdn") {
						fmt.Println(a.Val)
						tempInt := strconv.Itoa(i)
						i++
						DownloadFile("./temp/"+tempInt+".jpg", a.Val)
					}
					//break
				}
			}
			//fmt.Println("=")
		}
	}

	// 결과 출력
	fmt.Println("=======")
	/* bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
	*/
}
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
