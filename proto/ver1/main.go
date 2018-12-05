package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	//response, err := http.Get("https://www.devdungeon.com")
	response, err := http.NewRequest("GET", "https://kissme2145.tistory.com/1418?category=634440", nil)
	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	response.Header.Add("User-Agent", "Crawler")
	//response, err := http.Get("http://localhost:8090/getMyBook")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		panic(err)

	}

	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print image URLs
	i := 0
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			fmt.Println(imgSrc) // 굳이 보여줄 필요는...
			tempInt := strconv.Itoa(i)
			i++
			DownloadFile("./temp/"+tempInt+".jpg", imgSrc)
		}
	})
}

func DownloadFile(filepath string, url string) error {

	//strings.Split(filepath, "/")[0]

	// Create the file
	//strins.filepath.IsDir()
	// 파일 패스는 depth 가 여러개 들어 갈 수 있음
	//os.IsDir()
	filepathOnlyPath := strings.Split(filepath, "/")[:3]
	file, err := os.Open(filepathOnlyPath)
	if err != nil {
		// handle the error and return
	}
	defer file.Close()

	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
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
