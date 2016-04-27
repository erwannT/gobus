package main

import (
	"database/sql"
	"encoding/csv"
	"gobus/utils"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

func runParseRoute(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	file, err := os.Open("/home/erwann/Documents/GTFS-20151218/routes.txt")
	utils.Check(err)
	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	utils.Check(err)
	insertRoutes(rows)
}

func insertRoutes(rows [][]string) {
	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	utils.Check(err)
	defer db.Close()

	utils.Check(err)
	tx, err := db.Begin()
	utils.Check(err)
	prepareStmt, err := tx.Prepare("INSERT INTO route(" +
		"route_id, agency_id, route_short_name, route_long_name, route_desc," +
		"route_type, route_url, route_color, route_text_color)" +
		"VALUES ($1, $2, $3, $4, $5," +
		"$6, $7, $8, $9);")
	utils.Check(err)
	for id, row := range rows {
		if id != 0 {
			log.Println(row)
			_, err = prepareStmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8])
			utils.Check(err)
		}
	}

	err = tx.Commit()
	utils.Check(err)
	prepareStmt.Close()
}
