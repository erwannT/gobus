package web

import (
	"encoding/json"
	"fmt"
	"gobus/calculator"
	"gobus/calculator/gtfsDao"
	"gobus/calculator/model"
	"gobus/utils"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func Start() {

	http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("/home/erwann/tmp/gobus-web"))))
	http.HandleFunc("/findroute", api_findRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func api_findRoute(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	from := strings.Split(params.Get("from"), ",")
	to := strings.Split(params.Get("to"), ",")
	// startDate := params.Get("time")
	// if startDate == "now" {
	startDate := time.Now().Format(utils.PatternhhmmForm)
	// }

	var dao *gtfsDao.GtfsDao
	dao = gtfsDao.NewGtfsDao()
	defer dao.Close()

	// determiner le point de d'arrivée
	endPosition := model.Position{PositionLong: to[0], PositionLat: to[1]}
	arrivalStops := dao.FindStop(endPosition, 50)

	// déterminer le point de depart
	startPosition := model.Position{PositionLong: from[0], PositionLat: from[1]}

	c := make(chan model.ReponseProcessChan)
	var wg sync.WaitGroup
	wg.Add(1)
	go calculator.ProcessChan(dao, 1, startDate, startPosition, model.Trip{}, arrivalStops, c, &wg)
	response := <-c

	//response := calculator.Process(dao, 1, startDate, startPosition, model.Trip{}, arrivalStops)

	directions := response.Directions

	journeys := model.TJourneys{}

	for _, direction := range directions {
		journeys.To(model.TJourney{}, *direction)
	}

	w.Header().Set("Content-Type", "application/json")
	resJSON, _ := json.Marshal(journeys)
	fmt.Fprintf(w, string(resJSON))
}
