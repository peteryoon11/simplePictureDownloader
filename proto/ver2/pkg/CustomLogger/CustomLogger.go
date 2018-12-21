package CustomLogger

import (
	"fmt"
	"io"
	"log"
	"os"
)

func TestFunc() {

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
		fmt.Println(err)
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
	workerRecorder = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	workerRecorder.SetOutput(multiWriter)
	fmt.Println("test")

}
