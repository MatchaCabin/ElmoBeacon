package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func InitLogger() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05 Z07:00"
	logFile, err := os.OpenFile("ElmoBeacon.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.Logger = log.Output(multiWriter)
}
