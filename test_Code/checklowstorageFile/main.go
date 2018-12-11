package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	/*  original web source


	 */

	/*
			Date →Mon, 10 Dec 2018 08:42:20 GMT
		Server →PWS/8.3.2.1
		X-Px →ms h0-s1336.p61-icn ( h0-s1317.p61-icn), ht h0-s1317.p61-icn.cdngp.net
		Age →170346
		Cache-Control →max-age=172800
		Expires →Mon, 10 Dec 2018 09:23:14 GMT
		Accept-Ranges →bytes
		Content-Length →387977
		Content-Type →image/jpeg
		Last-Modified →Tue, 10 Apr 2018 01:56:45 GMT
		Connection →keep-alive

		추출한 이미지 소스의 경로를 get 으로 요청하면 위와 같은 헤더값을  던져 준다.
		여기서 Content-Type 과
		Content-Length
		Accept_Ranges 를 가지고
		확장자 명이랑 받지 않을 파일을 지정 하도록 하자.
		받지 않을때는 마지막에 받지 않은 이유  기타 이미지, 작은 이미지 등을 출력해 주도록 하자.

	*/
	//fileUrl := "https://golangcode.com/images/avatar.jpg"
	//fileUrl := "https://t1.daumcdn.net/cfile/tistory/99120B3F5ACC19DE0E"
	fileUrl := "https://t1.daumcdn.net/cfile/tistory/994700375BFA409B2A" // gif

	err := DownloadFile("./temp/avatar3.gif", fileUrl)
	if err != nil {
		panic(err)
	}

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
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
	fmt.Println(resp.Header)
	for key, item := range resp.Header {
		fmt.Println("key = ", key, " item = ", item)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
