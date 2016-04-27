package model

type TPos struct {
	Lon float64
	Lat float64
}

type TTimeStop struct {
	Time     string
	StopName string
}

type TTrip struct {
	Line      string
	Direction string
	Departure TTimeStop
	Arrival   TTimeStop
}

type TJourney struct {
	Steps []TTrip
}

type TJourneys []TJourney

func (journeys *TJourneys) To(current TJourney, direction Direction) {

	trip := TTrip{}
	trip.Direction = direction.Headsign
	trip.Line = direction.Route
	trip.Departure = TTimeStop{}
	trip.Departure.Time = direction.StartPoint.Departuretime
	trip.Departure.StopName = direction.StartPoint.StopName
	trip.Arrival = TTimeStop{}
	trip.Arrival.Time = direction.EndPoint.Arrivaltime
	trip.Arrival.StopName = direction.EndPoint.StopName

	current.Steps = append(current.Steps, trip)

	if direction.NextDirections == nil {
		*journeys = append(*journeys, current)
	} else {

		for _, v := range direction.NextDirections {
			journeys.To(current, *v)

		}
	}

}
