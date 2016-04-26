package main

import (
	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/lib/pq"
)

func runParseTrip() {

	file, err := os.Open("/home/erwann/Documents/GTFS-20151218/trips.txt")
	check(err)

	csvReader := csv.NewReader(file)

	rows, err := csvReader.ReadAll()
	check(err)

	insertTrips(rows)

}

func insertTrips(rows [][]string) {
	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	check(err)
	defer db.Close()

	prepareStmt, err := db.Prepare("INSERT INTO trip(" +
		"route_id, service_id, trip_id, trip_headsign, trip_short_name, " +
		"direction_id, block_id, shape_id, wheelchair_accessible, bikes_allowed)" +
		"VALUES ($1, $2, $3, $4, $5, " +
		"$6, $7, $8, $9, $10)")

	check(err)
	tx, err := db.Begin()
	check(err)
	stmt := tx.Stmt(prepareStmt)
	for id, row := range rows {
		if id != 0 {
			_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9])
			check(err)
		}
	}
	stmt.Close()
	err = tx.Commit()
	check(err)
	prepareStmt.Close()

}
