package CustomLogger

import (
	"fmt"
	"io"
	"log"
	"os"
)

//var fpLog *os.File

func LoggerEnd(fpLog *os.File) {
	fpLog.Close() // 우선 여기서 종료 하는데.. 추후에 문제가 생기지 않을까? 아닌가?
}

func LoggerAgent(loggerLocate string, workerRecorder *log.Logger) (*log.Logger, *os.File) {
	var (
		err error
		//workerRecorder_return *log.Logger
	)

	fpLog, err := os.OpenFile("./"+loggerLocate+"/logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		//panic(err)
		fmt.Println(err)
		log.Println(err)
	}

	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)
	//log = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	workerRecorder = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	workerRecorder.SetOutput(multiWriter)
	fmt.Println("test")
	//return nil, workerRecorder
	return workerRecorder, fpLog
}
