package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var myLogger *log.Logger

func main() {
	LoggingCore()
	p := fmt.Println
	now := time.Now()
	p(now)
	//time.Sleep(100 * time.Millisecond)
	time.Sleep(2 * time.Second)

	diff := now.Sub(time.Now())
	p(diff)
	p(diff * (-1))
	//TestTimePrint()
}

func statusUpdate() string { return "" }

func TestCode() {
	c := time.Tick(1 * time.Minute)
	for now := range c {
		fmt.Printf("%v %s\n", now, statusUpdate())
	}
}

func LoggingCore() {
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
	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	run()
	myLogger.Println("End of Program")
}
func TestTimePrint() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("=======================")
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p("=======================")
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
func run() {
	myLogger.Print("Test mylogger")
	log.Print("Test")
}
