package utils

import (
	"log"
	"time"
)

const (
	PatternhhmmForm = "15:04:05"
)

func Add(currentTime string, addMinute int) (timeAdded string, parseErr error) {

	sh, err := time.Parse(PatternhhmmForm, currentTime)
	if err != nil {
		parseErr = err
		return
	}
	eh := sh.Add(time.Duration(addMinute) * time.Minute)
	timeAdded = eh.Format(PatternhhmmForm)
	return
}

func Check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
