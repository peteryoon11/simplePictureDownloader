// http_get101.go
//
// get the contents of a web page with given URL
//
// for imported package info see ...
// http://golang.org/pkg/fmt/
// http://golang.org/pkg/io/ioutil/
// http://golang.org/pkg/net/http/
//
// tested with Go version 1.4.2   by vegaseat  28apr2015
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	/*


	 */
	improvGetWebSourceFunc()
	//originGetFunc()
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
		recover()
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)

}
func originGetFunc() {
	//url := "http://tour.golang.org/welcome/1"
	//url := "https://kissme2145.tistory.com/1287"
	url := "https://kissme2145.tistory.com/1418?category=634440"
	fmt.Printf("HTML code of %s ...\n", url)

	resp, err := http.Get(url)

	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
}
