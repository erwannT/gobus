package calculator

import (
	"gobus/calculator/gtfsDao"
	"gobus/calculator/model"
	"gobus/utils"
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
func Process(dao *gtfsDao.GtfsDao, n int, startTime string, currentPosition model.Position, sourceTrip model.Trip, arrivalStops []model.Stop) (response model.ReponseProcessChan /*directions []*model.Direction*/) {

	startHour := startTime
	endHour, err := utils.Add(startTime, 10)
	utils.Check(err)

	currentPositionTime := model.PositionTime{Position: model.Position{PositionLong: currentPosition.PositionLong, PositionLat: currentPosition.PositionLat}, StartHour: startHour, EndHour: endHour}

	trips := dao.FindDirections(currentPositionTime, 200, sourceTrip.RouteId)

	var directions []*model.Direction

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
		for _, trip := range trips {

			if n+1 <= differentBusMax {

				currentPosition := model.Position{PositionLong: trip.EndPoint.Stop.Xpos, PositionLat: trip.EndPoint.Stop.Ypos}
				startTime = trip.EndPoint.Arrivaltime

				response = Process(dao, n+1, startTime, currentPosition, trip, arrivalStops)
				if len(response.Directions) != 0 {
					var direction = &model.Direction{Trip: response.SourceTrip, NextDirections: response.Directions}
					directions = append(directions, direction)
				}
			}
		}

	}
	return model.ReponseProcessChan{SourceTrip: sourceTrip, Directions: directions}

}
