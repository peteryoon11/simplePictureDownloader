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
	"time"

	"github.com/PuerkitoBio/goquery"
)

//var log *log.Logger //
var fpLog *os.File

func main() {
	// Make HTTP request
	//response, err := http.Get("https://www.devdungeon.com")
	//fmt.Println(os.Args[1:])
	var (
		webpageAddress string
		filepath       string
		identify       string
		loggerLocate   string
	)
	startTime := time.Now() // 처음부터 끝까지 걸린 시간을 측정 하기 위한 시작시간 체크
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
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "logger") {
			// logger 파일의 위치
			loggerLocate = strings.Split(item, "=")[1]
		}
	}
	//fmt.Println("web page!!", webpageAddress)
	//fmt.Println("path!!", filepath)
	// 필요한 부분은 웹페이지 하나 일단
	//	ProcessCore()
	LoggerAgent(loggerLocate)
	ProcessCore(webpageAddress, filepath, identify, loggerLocate, startTime)
	LoggerEnd()

}
func LoggerEnd() {
	fpLog.Close() // 우선 여기서 종료 하는데.. 추후에 문제가 생기지 않을까? 아닌가?
}
func LoggerAgent(loggerLocate string) {
	var (
		err error
	)
	if len(loggerLocate) == 0 {
		loggerLocate = "logs"
	}
	// 나중에 여기에 경로가 있는지 판별하고 만드는 부분 모듈화 해서 호출 하기
	// 지금은 테스트 니까 그냥 만들어 두고 쓰자.
	// 로그파일이 하루 지났으면 이전 파일을 날짜붙여서 백업하고 새로 만들어서
	// 로깅을 시작 할까?

	fpLog, err = os.OpenFile("./"+loggerLocate+"/logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		//panic(err)
		//fmt.Println(err)
		log.Println(err)
	}

	//defer fpLog.Close()
	/*
		이슈 등록 해야지
			이 부분이 있어서 함수 가 종료되고 나서 log 객체를 닫아 버려서 기록이 안됨..
			닫긴 닫아야 하지 않을까? 언제 닫지..
			구성을 바꿔야 할거 같다. 메인에서 시작했다가 닫는 부분으로
			이전에 다른 곳에서는 log init 이라는 걸로 시작 했다가 닫는 부분이 있었으니까
			그 부분을 참고해서 만들어 보자.
	*/
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)
	//log = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	fmt.Println("test")
}
func ProcessCore(webpage string, filepath string, identify string, loggerLocate string, startTime time.Time) {

	if len(webpage) == 0 {
		webpage = "https://kissme2145.tistory.com/1418?category=634440"
		//webpage = "https://comic.naver.com/webtoon/detail.nhn?titleId=675554&no=683"
		// 나중에 네이버 웹툰 페이지도 추가해 보자.
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
		//fmt.Println(err)
		log.Println(err)
	}

	log.Println("Start SimlePicDownloader!!!!!!!!!!!")
	log.Println("from ", webpage)
	log.Println("to ", filepath) // 여기는 절대 경로로 보여주는걸로 바꿔주자.
	log.Println("filename base is  ", identify)
	log.Println("logfile will locate  ", loggerLocate)

	// 로그 파일 위치를 보여주자.

	//필요시 헤더 추가 가능
	//response.Header.Add("Referer", "Crawler") // 이전 사이트의 정보?
	response.Header.Add("User-Agent", "Crawler")
	//response, err := http.Get("http://localhost:8090/getMyBook")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		//panic(err)
		//fmt.Println(err)
		//	fmt.Println("Check WAN connect")
		log.Println("Check WAN connect")
		return
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
	diff := startTime.Sub(time.Now())
	/* fmt.Println("total spend time is ", (diff * (-1)))
	fmt.Println("total download image is ", (i + 1)) */
	log.Println("total spend time is ", (diff * (-1)))
	log.Println("total download image is ", (i + 1))
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
	/*
		파일을 만드려고 하면 그때 파일이 없어서 or 패스가 없어서 오류가 생기면
		해당 시점에 그 경로를 만드는 코드를 만들자.
		이거는 지금 이미지 저장 부분이나 로그 파일을 만드는 부분에 추가 해서
		쓰면 코드 재 사용 성이나 나중에 수정 할때나 편리 할듯
	*/
	filepathOnlyPath, _ := path.Split(filepath)
	if count == 0 {

		if runtime.GOOS == "linux" || runtime.GOOS == "darwin" { // also can be specified to FreeBSD

			log.Println("Unix/Linux or Mac OS type OS detected")
			if _, err := os.Stat(filepathOnlyPath); os.IsNotExist(err) {

				err := os.Mkdir(filepathOnlyPath, 0755)
				if err != nil {
					log.Println(err)

				}
			} else {
				log.Println("filepath is exist")

			}
		}
		if runtime.GOOS == "windows" {

			log.Println("Windows OS detected")
			if _, err := os.Stat(filepathOnlyPath); !os.IsNotExist(err) {

				err := os.Mkdir(filepathOnlyPath, 0755)
				if err != nil {

					log.Println(err)
				}
			} else {
				log.Println("filepath is exist")

			}
		}
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, count
	}
	defer resp.Body.Close()

	Filesize, err := strconv.Atoi(resp.Header["Content-Length"][0])
	if nil != err {
		log.Println(err)
		return nil, count
		//fmt.Println(err)
	}

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
			//fmt.Println(err)
			log.Println(err)

			return nil, count
		}
		/* fmt.Println("url = ", url)
		fmt.Println("filesize = ", Filesize/1000, " kbytes")
		*/
		//log.Println("url = ", url)
		//log.Println("filesize = ", Filesize/1000, " kbytes")
		//	fmt.Println("25000 bytes 이상 / 25kbytes")
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Println(err)
			return nil, count
		}
		count++
		defer out.Close()
	} else {
		//fmt.Println("25000 bytes 미만 / 24kbytes")
		log.Println("25000 bytes 미만 / 24kbytes")
		// 다운로드 받지 않은 url 및 사이즈 보여주자.
		//fmt.Println("url = ", url)
		log.Println("url = ", url)
		//fmt.Println("filesize = ", Filesize/1000, " kbytes")
		log.Println("filesize = ", Filesize/1000, " kbytes")
		//	return nil, count
	}
	//	log.Println("현재 count!! ", count)
	//fmt.Println("현재 count!! ", count)
	// Write the body to file
	/* 	_, err = io.Copy(out, resp.Body)
	   	if err != nil {
	   		return nil, count
	   	}
	   	count++ */

	return nil, count
}
