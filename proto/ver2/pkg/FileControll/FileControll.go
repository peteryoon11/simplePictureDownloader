package FileControll

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

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
func CheckOSAndMakeFile(filepath string, workerRecorder *log.Logger) {

	filepathOnlyPath, _ := path.Split(filepath)
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" { // also can be specified to FreeBSD

		workerRecorder.Println("Unix/Linux or Mac OS type OS detected")
		if _, err := os.Stat(filepathOnlyPath); os.IsNotExist(err) {

			err := os.Mkdir(filepathOnlyPath, 0755)
			if err != nil {
				workerRecorder.Println(err)

			}
		} else {
			workerRecorder.Println("filepath is exist")

		}
	}
	if runtime.GOOS == "windows" {

		workerRecorder.Println("Windows OS detected")
		if _, err := os.Stat(filepathOnlyPath); !os.IsNotExist(err) {

			err := os.Mkdir(filepathOnlyPath, 0755)
			if err != nil {

				workerRecorder.Println(err)
			}
		} else {
			workerRecorder.Println("filepath is exist")

		}
	}

}
func DownloadFile(filepath string, url string, count int, workerRecorder *log.Logger) (error, int) {

	/*
		파일을 만드려고 하면 그때 파일이 없어서 or 패스가 없어서 오류가 생기면
		해당 시점에 그 경로를 만드는 코드를 만들자.
		이거는 지금 이미지 저장 부분이나 로그 파일을 만드는 부분에 추가 해서
		쓰면 코드 재 사용 성이나 나중에 수정 할때나 편리 할듯
	*/

	//filepathOnlyPath, _ := path.Split(filepath)
	if count == 0 {
		CheckOSAndMakeFile(filepath, workerRecorder)

	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		workerRecorder.Println(err)
		return nil, count
	}
	defer resp.Body.Close()

	Filesize, err := strconv.Atoi(resp.Header["Content-Length"][0])
	if nil != err {
		workerRecorder.Println(err)
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
			workerRecorder.Println(err)

			return nil, count
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			workerRecorder.Println(err)
			return nil, count
		}
		count++
		defer out.Close()
	} else {

		workerRecorder.Println("25000 bytes 미만 / 24kbytes")
		// 다운로드 받지 않은 url 및 사이즈 보여주자.

		workerRecorder.Println("url = ", url)

		workerRecorder.Println("filesize = ", Filesize/1000, " kbytes")

	}

	return nil, count
}
