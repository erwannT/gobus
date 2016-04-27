package main

import (
	"database/sql"
	"encoding/csv"
	"gobus/utils"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

func runParseTrip(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	file, err := os.Open("/home/erwann/Documents/GTFS-20151218/trips.txt")
	utils.Check(err)

	csvReader := csv.NewReader(file)

	rows, err := csvReader.ReadAll()
	utils.Check(err)

	insertTrips(rows)

}

func insertTrips(rows [][]string) {
	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	utils.Check(err)
	defer db.Close()

	utils.Check(err)
	tx, err := db.Begin()
	utils.Check(err)

	prepareStmt, err := db.Prepare("INSERT INTO trip(" +
		"route_id, service_id, trip_id, trip_headsign, trip_short_name, " +
		"direction_id, block_id, shape_id, wheelchair_accessible, bikes_allowed)" +
		"VALUES ($1, $2, $3, $4, $5, " +
		"$6, $7, $8, $9, $10)")
	utils.Check(err)
	for id, row := range rows {
		if id != 0 {
			_, err = prepareStmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9])
			utils.Check(err)
		}
	}
	err = tx.Commit()
	utils.Check(err)
	prepareStmt.Close()

}
