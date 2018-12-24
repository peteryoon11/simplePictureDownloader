package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"../pkg/CustomLogger"
	"../pkg/FileControll"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	// 프로그램 시작 시에 값을 초기화 하는 부분
	initFunc(os.Args[1:])

}
func initFunc(startWord []string) {
	var (
		webpageAddress string
		filepath       string
		identify       string
		loggerLocate   string
		workerRecorder *log.Logger
		fpLog          *os.File
	)
	startTime := time.Now() // 처음부터 끝까지 걸린 시간을 측정 하기 위한 시작시간 체크

	for _, item := range os.Args[1:] {
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

	if len(webpageAddress) == 0 {
		webpageAddress = "https://kissme2145.tistory.com/1418?category=634440"
		//webpage = "https://comic.naver.com/webtoon/detail.nhn?titleId=675554&no=683"
		// 나중에 네이버 웹툰 페이지도 추가해 보자.
	}
	if len(filepath) == 0 {
		filepath = "temp"
	}
	if len(identify) == 0 {
		identify = "image"
	}
	if len(loggerLocate) == 0 {
		loggerLocate = "logs"
	}

	workerRecorder, fpLog = CustomLogger.LoggerAgent(loggerLocate, workerRecorder)

	workerRecorder.Println("Start SimlePicDownloader!!!!!!!!!!!")
	workerRecorder.Println("from ", webpageAddress)
	workerRecorder.Println("to ", filepath) // 여기는 절대 경로로 보여주는걸로 바꿔주자.
	workerRecorder.Println("filename base is  ", identify)
	workerRecorder.Println("logfile will locate  ", loggerLocate)
	ProcessCore(webpageAddress, filepath, identify, loggerLocate, startTime, workerRecorder)

	CustomLogger.LoggerEnd(fpLog)

}

func ProcessCore(webpage string, filepath string, identify string, loggerLocate string, startTime time.Time, workerRecorder *log.Logger) {

	response, err := http.NewRequest("GET", webpage, nil)
	if err != nil {
		workerRecorder.Println(err)
	}

	response.Header.Add("User-Agent", "Crawler")

	if err != nil {
		workerRecorder.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(response)
	if err != nil {
		workerRecorder.Println("Check WAN connect")
		return
	}

	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		workerRecorder.Fatal("Error loading HTTP response body. ", err)
	}

	i := 0
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			tempInt := FileControll.DisplayNumberSort(i)
			tempFilename := identify + tempInt
			_, i = FileControll.DownloadFile("./"+filepath+"/"+tempFilename, imgSrc, i, workerRecorder)
		}
	})
	diff := startTime.Sub(time.Now())
	workerRecorder.Println("total spend time is ", (diff * (-1)))
	workerRecorder.Println("total download image is ", (i + 1))
}
