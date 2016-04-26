package main

import (
	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/lib/pq"
)

func runParseStopTime() {

	file, err := os.Open("/home/erwann/Documents/GTFS-20151218/stop_times.txt")
	check(err)

	csvReader := csv.NewReader(file)

	rows, err := csvReader.ReadAll()
	check(err)

	insertStopTimes(rows)

}

func insertStopTimes(rows [][]string) {
	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	check(err)
	defer db.Close()

	prepareStmt, err := db.Prepare("INSERT INTO stop_time(" +
		"trip_id, arrival_time, departure_time, stop_id, stop_sequence, " +
		"stop_headsign, pickup_type, drop_off_type, shape_dist_traveled)" +
		"VALUES ($1, $2, $3, $4, $5, " +
		"$6, $7, $8, $9)")

	check(err)
	tx, err := db.Begin()
	check(err)
	stmt := tx.Stmt(prepareStmt)
	for id, row := range rows {
		if id != 0 {
			_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8])
			check(err)
		}
	}
	stmt.Close()
	err = tx.Commit()
	check(err)
	prepareStmt.Close()

}
