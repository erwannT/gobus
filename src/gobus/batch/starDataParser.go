package main

// import "os"
import (
	"log"
	"sync"
)

func startParser() {

	log.Println("Demarrage du batch")

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	runParseStop(&waitGroup)
	waitGroup.Add(1)
	runParseRoute(&waitGroup)
	waitGroup.Add(1)
	runParseCalendar(&waitGroup)
	waitGroup.Add(1)
	runParseStopTime(&waitGroup)
	waitGroup.Add(1)
	runParseTrip(&waitGroup)

	log.Println("En attente de la fin du batch")
	waitGroup.Wait()

}
