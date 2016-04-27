/*
Package web implement un serveur simple
*/
package web

import (
	"encoding/json"
	"fmt"
	"gobus/gtfs"
	"gobus/gtfs/dao"
	"gobus/gtfs/model"
	"gobus/utils"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

/*Start demarre le serveur web*/
func Start() {

	http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("/home/erwann/tmp/gobus-web"))))
	http.HandleFunc("/findroute", apiFindRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func apiFindRoute(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	from := strings.Split(params.Get("from"), ",")
	to := strings.Split(params.Get("to"), ",")
	// startDate := params.Get("time")
	// if startDate == "now" {
	startDate := time.Now().Format(utils.PatternhhmmForm)
	// }

	dao := dao.NewGtfsDao()
	defer dao.Close()

	// determiner le point de d'arrivée
	endPosition := model.Position{PositionLong: to[0], PositionLat: to[1]}
	arrivalStops := dao.FindStop(endPosition, 50)

	// déterminer le point de depart
	startPosition := model.Position{PositionLong: from[0], PositionLat: from[1]}

	c := make(chan model.ReponseProcess)
	var wg sync.WaitGroup
	wg.Add(1)
	go gtfs.ProcessChan(dao, 1, startDate, startPosition, model.Trip{}, arrivalStops, c, &wg)
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
