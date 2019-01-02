package StructureModule

import (
	"log"
	"time"
)

type Repository struct {
	ArchiveURL string
}

type ProcessCoreMandantory struct {
	webpageAddress string
	filepath       string
	identify       string
	loggerLocate   string
	startTime      time.Time
	workerRecorder *log.Logger
}
