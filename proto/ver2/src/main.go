package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	//response, err := http.Get("https://www.devdungeon.com")
	//fmt.Println(os.Args[1:])
	var (
		webpageAddress string
		filepath       string
		identify       string
	)
	for _, item := range os.Args[1:] {
		//fmt.Println(item)
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "site") {
			webpageAddress = strings.Split(item, "=")[1]
		}
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "path") {
			filepath = strings.Split(item, "=")[1]
		}
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "identi") {
			identify = strings.Split(item, "=")[1]
		}
	}
	//fmt.Println("web page!!", webpageAddress)
	//fmt.Println("path!!", filepath)
	// 필요한 부분은 웹페이지 하나 일단
	//	ProcessCore()
	ProcessCore(webpageAddress, filepath, identify)

}
func ProcessCore(webpage string, filepath string, identify string) {

	//response, err := http.NewRequest("GET", "https://kissme2145.tistory.com/1418?category=634440", nil)
	if len(webpage) == 0 {
		webpage = "https://kissme2145.tistory.com/1418?category=634440"
	}
	if len(filepath) == 0 {
		filepath = "temp"
	}
	if len(identify) == 0 {
		identify = "image"
	}
	response, err := http.NewRequest("GET", webpage, nil)
	if err != nil {
		//panic(err)
		fmt.Println(err)
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
		//panic(err)
		fmt.Println(err)
		fmt.Println("Check wan connect")

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
			//	fmt.Println(imgSrc) // 굳이 보여줄 필요는...
			//tempInt := strconv.Itoa(i)
			tempInt := DisplayNumberSort(i)
			tempFilename := identify + tempInt
			//i++
			//DownloadFile("./temp/"+tempInt+".jpg", imgSrc)
			//_, i = DownloadFile("./"+filepath+"/"+tempInt+".jpg", imgSrc, i)
			_, i = DownloadFile("./"+filepath+"/"+tempFilename, imgSrc, i)
		}
	})
	fmt.Println("total download image is ", (i + 1))
}
func DisplayNumberSort(givennumber int) string {
	// 000 자리로 나오게 설정
	// ex) 001 002 ~~ 010 011 ~~ 100 101 ~~ 201 202
	var result string
	if givennumber < 10 {
		tempInt := strconv.Itoa(givennumber)
		result = "00" + tempInt
	} else if (givennumber >= 10) && (givennumber < 100) {
		tempInt := strconv.Itoa(givennumber)
		result = "0" + tempInt
	} else {
		result = strconv.Itoa(givennumber)
	}
	return result

}
func DownloadFile(filepath string, url string, count int) (error, int) {

	//strings.Split(filepath, "/")[0]

	// Create the file
	//strins.filepath.IsDir()
	// 파일 패스는 depth 가 여러개 들어 갈 수 있음
	//os.IsDir()
	filepathOnlyPath, _ := path.Split(filepath)
	if count == 0 {

		if runtime.GOOS == "linux" || runtime.GOOS == "darwin" { // also can be specified to FreeBSD
			fmt.Println("Unix/Linux or Mac OS type OS detected")
			if _, err := os.Stat(filepathOnlyPath); os.IsNotExist(err) {

				err := os.Mkdir(filepathOnlyPath, 0755)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("filepath is exist")
			}
		}
		if runtime.GOOS == "windows" {
			fmt.Println("Windows OS detected")
			if _, err := os.Stat(filepathOnlyPath); !os.IsNotExist(err) {

				err := os.Mkdir(filepathOnlyPath, 0755)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("filepath is exist")
			}
		}
	}
	/* 	if _, err := os.Stat(filepathOnlyPath); !os.IsNotExist(err) {
	   		//!os.IsNotExist(err) for window
	   		// os.IsNotExist(err)  for mac
	   		// path/to/whatever does not exist
	   		//fmt.Println("filepathOnlyPath = ", filepathOnlyPath)
	   		// 이렇게 여러번 확인 할 필요가 있나.. 싶은데.. 나중에 다시 체크 하자.
	   		err := os.Mkdir(filepathOnlyPath, 0755)
	   		if err != nil {
	   			fmt.Println(err)
	   		}
	   	} else {
	   		fmt.Println("filepath is exist")
	   	} */

	/* 	file, err := os.Open(filepathOnlyPath)
	   	if err != nil {
	   		// handle the error and return
	   		fmt.Println(err)
	   	} */
	/* defer file.Close()
	fi, err := flie.Stat() */
	/* if err != nil {
		// handle the error and return
		fmt.Println(err)
	}
	if !file.IsDir() {
		os.Mkdir(filepathOnlyPath, 0755)
	}
	*/

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, count
	}
	defer resp.Body.Close()

	Filesize, err := strconv.Atoi(resp.Header["Content-Length"][0])
	if nil != err {
		fmt.Println(err)
	}

	fmt.Println(strings.Split(resp.Header["Content-Type"][0], "/")[1])
	//	out, err := os.Create( filepath+"."+ resp.Header["Content-Type"][0], "/")[1]))
	// out, err := os.Create(filepath + "." + strings.Split(resp.Header["Content-Type"][0], "/")[1])

	if Filesize > 24999 {
		/*
			기준은 25kb 이하만 디폴트로 다운로드 하지 않을거임
			1000 bytes = 1 kbytes
			25000 bytes = 25 kbytes
		*/
		//out, err := os.Create(filepath)
		out, err := os.Create(filepath + "." + strings.Split(resp.Header["Content-Type"][0], "/")[1])
		if err != nil {
			//fmt.Println("create")
			fmt.Println(err)
			return nil, count
		}
		fmt.Println("url = ", url)
		fmt.Println("filesize = ", Filesize/1000, " kbytes")
		//	fmt.Println("25000 bytes 이상 / 25kbytes")
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return nil, count
		}
		count++
		defer out.Close()
	} else {
		fmt.Println("25000 bytes 미만 / 24kbytes")
		// 다운로드 받지 않은 url 및 사이즈 보여주자.
		fmt.Println("url = ", url)
		fmt.Println("filesize = ", Filesize/1000, " kbytes")
		//	return nil, count
	}
	fmt.Println("현재 count!! ", count)
	// Write the body to file
	/* 	_, err = io.Copy(out, resp.Body)
	   	if err != nil {
	   		return nil, count
	   	}
	   	count++ */

	return nil, count
}
