package main

import (
	"io"
	"log"
	"os"
)

var myLogger *log.Logger

func main() {
	// 로그파일 오픈
	/* 	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	   	if err != nil {
	   		panic(err)
	   	}
	   	defer fpLog.Close()

	   	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	   	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	   	log.SetOutput(multiWriter)

	   	//....
	   	run()

		   myLogger.Println("End of Program") */

	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	// 파일과 화면에 같이 출력하기 위해 MultiWriter 생성
	multiWriter := io.MultiWriter(fpLog, os.Stdout)
	log.SetOutput(multiWriter)

	run()
	log.Println("End of Program")
}

func run() {
	//myLogger.Print("Test")
	log.Print("Test")
}
