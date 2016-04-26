package main

import (
	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/lib/pq"
)

func runParseStop() {

	file, err := os.Open("/home/erwann/Documents/GTFS-20151218/stops.txt")
	check(err)
	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	check(err)
	insertStops(rows)

}

func insertStops(rows [][]string) {
	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	check(err)
	defer db.Close()

	prepareStmt, err := db.Prepare("INSERT INTO stop( stop_id, stop_code, stop_name, stop_desc, stop_coord, zone_id, stop_url, location_type, parent_station, stop_timezone, wheelchair_boarding)  " +
		"VALUES ($1, $2, $3, $4, ST_GeomFromText('POINT('|| $5 ||'' || $6 ||')', 4326)," +
		"$7, $8, $9, $10, $11, $12)")
	check(err)
	tx, err := db.Begin()
	check(err)
	stmt := tx.Stmt(prepareStmt)
	for id, row := range rows {
		if id != 0 {
			_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11])
			check(err)
		}
	}
	stmt.Close()
	err = tx.Commit()
	check(err)
	prepareStmt.Close()
}
