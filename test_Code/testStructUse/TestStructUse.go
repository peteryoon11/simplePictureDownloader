package main

import (
	"fmt"
)

func main() {
	var m map[string]string

	m = make(map[string]string)
	//추가 혹은 갱신
	m["901"] = "Apple"
	m["134"] = "Grape"
	m["777"] = "Tomato"
	fmt.Println(ProcessCoreMandantory{"test", "test", "test", "test"})
	for num, item := range m {
		fmt.Println(num, item)
	}
}

type ProcessCoreMandantory struct {
	webpageAddress string
	filepath       string
	identify       string
	loggerLocate   string
	//startTime      time.Time
	//	workerRecorder *log.Logger
}
