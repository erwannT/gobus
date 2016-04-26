package gtfsDao

import (
	"database/sql"
	"gobus/calculator/model"
	"log"

	_ "github.com/lib/pq"
)

type GtfsDao struct {
	db *sql.DB
}

func NewGtfsDao() *GtfsDao {
	var err error
	dao := new(GtfsDao)
	dao.db, err = sql.Open("postgres", "user=erwann password=erwann database=star")

	dao.db.SetMaxOpenConns(10)
	check(err)
	return dao
}

func (dao *GtfsDao) Close() {
	dao.db.Close()
}

/**
*	Liste les points d'arrets en fonction de la position geo et d'une distance
*
 */
func (dao *GtfsDao) FindStop(position model.Position, distance int) []model.Stop {

	var stops []model.Stop

	rows, err := dao.db.Query("SELECT stop_id, stop_name,ST_X(ST_AsText(stop_coord)),ST_Y(ST_AsText(stop_coord)) FROM stop  WHERE ST_DWithin(stop_coord, ST_GeographyFromText('SRID=4326;POINT('|| $1 ||'' || $2 ||')'), $3) order by stop_id", position.PositionLong, position.PositionLat, distance)
	check(err)
	defer rows.Close()
	for rows.Next() {
		var stop model.Stop
		err := rows.Scan(&stop.StopID, &stop.StopName, &stop.Xpos, &stop.Ypos)
		check(err)
		stops = append(stops, stop)

	}
	return stops
}

/**
* Recherche les voyages disponibles entre deux horaires pour un ensemble d'arrets
*
 */
func (dao *GtfsDao) FindDirections(position model.PositionTime, distance int, currentRouteID int) []model.Trip {

	request := "select " +
		" t1.trip_id," +
		" t1.trip_headsign," +
		" route.route_id, " +
		" route.route_short_name, " +

		" stop_stop.stop_id," +
		" stop_stop.stop_name," +
		" stoptime_stop.departure_time," +
		" stoptime_stop.arrival_time," +
		" ST_X(ST_AsText(stop_stop.stop_coord))," +
		" ST_Y(ST_AsText(stop_stop.stop_coord)) ," +

		" stop_start.stop_id," +
		" stop_start.stop_name as depart," +
		" stoptime_start.departure_time," +
		" stoptime_start.arrival_time," +
		" ST_X(ST_AsText(stop_start.stop_coord))," +
		" ST_Y(ST_AsText(stop_start.stop_coord))" +

		" from stop_time stoptime_stop" +
		" 	join stop stop_stop on stoptime_stop.stop_id = stop_stop.stop_id" +
		" 	join trip t1 on stoptime_stop.trip_id = t1.trip_id" +
		" 	join stop_time stoptime_start on t1.trip_id = stoptime_start.trip_id" +
		" 	join stop stop_start on stoptime_start.stop_id = stop_start.stop_id" +
		"		join route route on route.route_id = t1.route_id " +

		" WHERE ST_DWithin(stop_start.stop_coord, ST_GeographyFromText('SRID=4326;POINT('|| $1 ||'' || $2 ||')'), $5)" +
		" and stoptime_start.arrival_time::time > $3 and stoptime_start.arrival_time::time < $4 " +
		" and  t1.service_id in (3,16,17,18,19,20,21)" +
		" and stoptime_stop.stop_sequence::int > stoptime_start.stop_sequence::int" +
		" and t1.route_id <> $6" +
		" order by t1.trip_id, stoptime_stop.stop_sequence::int"

	var trips []model.Trip

	rows, err := dao.db.Query(request, position.PositionLong, position.PositionLat, position.StartHour, position.EndHour, distance, currentRouteID)
	check(err)
	defer rows.Close()

	for rows.Next() {
		var trip model.Trip

		err := rows.Scan(&trip.Tripid, &trip.Headsign, &trip.RouteId, &trip.Route,
			&trip.EndPoint.StopID, &trip.EndPoint.StopName, &trip.EndPoint.Departuretime, &trip.EndPoint.Arrivaltime, &trip.EndPoint.Xpos, &trip.EndPoint.Ypos,
			&trip.StartPoint.StopID, &trip.StartPoint.StopName, &trip.StartPoint.Departuretime, &trip.StartPoint.Arrivaltime, &trip.StartPoint.Xpos, &trip.StartPoint.Ypos,
		)
		check(err)
		trips = append(trips, trip)

	}
	return trips
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
