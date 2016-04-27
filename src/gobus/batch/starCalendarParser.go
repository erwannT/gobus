package main

import (
	"database/sql"
	"encoding/csv"
	"gobus/utils"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

func runParseCalendar(waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()

	file, err := os.Open("/home/erwann/Documents/GTFS-20151218/calendar.txt")
	utils.Check(err)
	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	utils.Check(err)
	insertCalendar(rows)
}

func insertCalendar(rows [][]string) {
	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	utils.Check(err)
	defer db.Close()

	utils.Check(err)
	tx, err := db.Begin()
	prepareStmt, err := tx.Prepare("INSERT INTO calendar(" +
		"service_id, monday, tuesday, wednesday, thursday, friday, saturday, " +
		"sunday, start_date, end_date)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, " +
		" $8, $9, $10);")
	utils.Check(err)
	for id, row := range rows {
		if id != 0 {
			_, err = prepareStmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9])
			utils.Check(err)
		}
	}
	err = tx.Commit()
	prepareStmt.Close()
	utils.Check(err)
}
