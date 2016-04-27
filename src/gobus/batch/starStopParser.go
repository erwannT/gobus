package main

import (
	"database/sql"
	"encoding/csv"
	"gobus/utils"
	"os"

	"sync"

	_ "github.com/lib/pq"
)

func runParseStop(waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	file, err := os.Open("/home/erwann/Documents/GTFS-20151218/stops.txt")
	utils.Check(err)
	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	utils.Check(err)
	insertStops(rows)

}

func insertStops(rows [][]string) {
	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	utils.Check(err)
	defer db.Close()
	utils.Check(err)
	tx, err := db.Begin()
	utils.Check(err)
	prepareStmt, err := tx.Prepare("INSERT INTO stop( stop_id, stop_code, stop_name, stop_desc, stop_coord, zone_id, stop_url, location_type, parent_station, stop_timezone, wheelchair_boarding)  " +
		"VALUES ($1, $2, $3, $4, ST_GeomFromText('POINT('|| $5 ||'' || $6 ||')', 4326)," +
		"$7, $8, $9, $10, $11, $12)")
	utils.Check(err)
	for id, row := range rows {
		if id != 0 {
			_, err = prepareStmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11])
			utils.Check(err)
		}
	}
	err = tx.Commit()
	utils.Check(err)
	prepareStmt.Close()
}
