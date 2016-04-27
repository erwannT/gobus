package model

type Stop struct {
	StopID   int
	StopName string
	Xpos     string
	Ypos     string
}

type StopTime struct {
	Stop
	Departuretime string
	Arrivaltime   string
}

type Trip struct {
	StartPoint StopTime
	EndPoint   StopTime

	Tripid   string
	Headsign string

	RouteId int
	Route   string
}

type Direction struct {
	Trip
	NextDirections []*Direction
}
