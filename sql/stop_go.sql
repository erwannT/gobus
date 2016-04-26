-- determine les routes possibles en fonction des trip et des services disponible (ici le lundi)
	select 
	route.route_id,
	route.route_short_name,
	t1.trip_id, 
	t1.trip_headsign,

	stop_stop.stop_id,
	stop_stop.stop_name,
	stoptime_stop.departure_time, 
	stoptime_stop.arrival_time,  
	ST_X(ST_AsText(stop_stop.stop_coord)),
	ST_Y(ST_AsText(stop_stop.stop_coord)) ,

	stop_start.stop_id, 
	stop_start.stop_name as depart,
	stoptime_start.departure_time, 
	stoptime_start.arrival_time,  
	ST_X(ST_AsText(stop_start.stop_coord)),
	ST_Y(ST_AsText(stop_start.stop_coord))

	from stop_time stoptime_stop 
		join stop stop_stop on stoptime_stop.stop_id = stop_stop.stop_id 
		join trip t1 on stoptime_stop.trip_id = t1.trip_id  
		join stop_time stoptime_start on t1.trip_id = stoptime_start.trip_id
		join stop stop_start on stoptime_start.stop_id = stop_start.stop_id
		join route route on route.route_id = t1.route_id
		    
	WHERE ST_DWithin(stop_start.stop_coord, ST_GeographyFromText('SRID=4326;POINT(48.109970 -1.679217)'), 200) 
	and stoptime_start.arrival_time::time > '14:05' and stoptime_start.arrival_time::time < '14:20:00'
	and  t1.service_id in (3,16,17,18,19,20,21)
	and stoptime_stop.stop_sequence::int > stoptime_start.stop_sequence::int 
	order by t1.trip_id, stoptime_stop.stop_sequence::int