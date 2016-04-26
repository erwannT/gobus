package gtfsDao

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func FindStop(long string, lat string, distance int) {

	log.Println("Find stop for long : " + long + " lat : " + lat)

	db, err := sql.Open("postgres", "user=erwann password=erwann database=star")
	defer db.Close()
	check(err)
	log.Println("connecte")

	rows, err := db.Query("SELECT stop_id,stop_name, stop_desc FROM stop  WHERE ST_DWithin(stop_coord, ST_GeographyFromText('SRID=4326;POINT(? ?)'), ?) ", long, lat, distance)
	check(err)
	defer rows.Close()
	for rows.Next() {
		var stop_name string
		err := rows.Scan(&stop_name)
		check(err)
		log.Println(stop_name)

	}

	/*prepareStmt, err := db.Prepare("INSERT INTO calendar(" +
	"service_id, monday, tuesday, wednesday, thursday, friday, saturday, " +
	"sunday, start_date, end_date)" +
	"VALUES ($1, $2, $3, $4, $5, $6, $7, " +
	" $8, $9, $10);")*/

}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
