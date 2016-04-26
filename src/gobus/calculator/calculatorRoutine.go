package calculator

import (
	"gobus/calculator/gtfsDao"
	"gobus/calculator/model"
	"gobus/utils"
	"sync"
)

/**
* n : niveau de bus
* chemin : liste des different  stopTime
* startTime : l'heure de departure
* currentPosition : position courante
* arrivalStops : liste des point d'arrivée possible
*
* return : retourne la nouvelle d'arrivée possible
 */
func ProcessChan(dao *gtfsDao.GtfsDao, n int, startTime string, currentPosition model.Position, sourceTrip model.Trip, arrivalStops []model.Stop, out chan model.ReponseProcessChan, finalWG *sync.WaitGroup) {

	defer finalWG.Done()

	var directions []*model.Direction

	var wg sync.WaitGroup

	startHour := startTime
	endHour, err := utils.Add(startTime, 10)
	utils.Check(err)

	currentPositionTime := model.PositionTime{Position: model.Position{PositionLong: currentPosition.PositionLong, PositionLat: currentPosition.PositionLat}, StartHour: startHour, EndHour: endHour}

	trips := dao.FindDirections(currentPositionTime, 200, sourceTrip.RouteId)

	found := false
	// détermine si un chemin existe
	for _, trip := range trips {
		for _, arrivalStop := range arrivalStops {
			if arrivalStop.StopID == trip.EndPoint.StopID {
				directions = append(directions, &model.Direction{Trip: trip, NextDirections: nil})
				found = true
			}
		}
	}

	if found == false {
		chans := make(chan model.ReponseProcessChan, len(trips))
		for _, trip := range trips {
			if n+1 <= differentBusMax {

				currentPosition := model.Position{PositionLong: trip.EndPoint.Stop.Xpos, PositionLat: trip.EndPoint.Stop.Ypos}
				startTime = trip.EndPoint.Arrivaltime

				wg.Add(1)
				go ProcessChan(dao, n+1, startTime, currentPosition, trip, arrivalStops, chans, &wg)
			}
		}
		wg.Wait()

		for {
			var quit = false
			select {
			case response := <-chans:
				if len(response.Directions) != 0 {
					directions = append(directions, &model.Direction{Trip: response.SourceTrip, NextDirections: response.Directions})
				}
			default:
				quit = true
			}
			if quit {
				break
			}
		}
	}
	out <- model.ReponseProcessChan{SourceTrip: sourceTrip, Directions: directions}
}
