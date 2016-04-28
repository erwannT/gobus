package model

type TJourney struct {
	Steps []Trip
}

type TJourneys []TJourney

func (journeys *TJourneys) To(current TJourney, direction Direction) {

	current.Steps = append(current.Steps, direction.Trip)

	if direction.NextDirections == nil {
		*journeys = append(*journeys, current)
	} else {

		for _, v := range direction.NextDirections {
			journeys.To(current, *v)

		}
	}
}
