package main

import (
	"fmt"
)

func main() {
	fmt.Println(ProcessCoreMandantory("test", "test", "test", "test"))
}

type ProcessCoreMandantory struct {
	webpageAddress string
	filepath       string
	identify       string
	loggerLocate   string
	//startTime      time.Time
	//	workerRecorder *log.Logger
}
