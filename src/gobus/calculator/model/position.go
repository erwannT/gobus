package model

type Position struct {
	PositionLong string
	PositionLat  string
}

type PositionTime struct {
	Position
	StartHour string
	EndHour   string
}
