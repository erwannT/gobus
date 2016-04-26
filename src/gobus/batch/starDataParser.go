package main

// import "os"
import (
	"log"
	"sync"
)

func startParser() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	runParseStop()
	waitGroup.Add(1)
	runParseRoute(&waitGroup)
	waitGroup.Add(1)
	runParseCalendar()
	waitGroup.Add(1)
	runParseStopTime()
	waitGroup.Add(1)
	runParseTrip()
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
