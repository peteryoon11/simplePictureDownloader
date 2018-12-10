package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	/*  original web source

	https://kissme2145.tistory.com/1418?category=634440
	여기서 소스를 추출하고 파싱해서 다운로드를 하도록 하자.
	1. get 등으로 소스를 가져온다. 이때 header 에 값이 필요 한거 같다. tistory 는 agent 가 비어 있는 부분에 대해서 403 을 주는 거 같다.
	2. 소스 중에서 img 태그를 가지고 있는 부분을 파싱 한다.
		a. 파싱 하는 부분에 대한 설계가 필요
	3. 최종 적으로 위의 부분을 가지고 바이너리 파일에 따른 사용자 편의를 위한 cli 만들기
		a. 예상 추후 변동 가능
		simpleDownloader -site=https://site.com -path=./twice/member -parsingType=img
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

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
