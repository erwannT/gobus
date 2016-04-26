 -- SELECT st_x(st_astext(stop_coord)), st_y(st_astext(stop_coord))
 -- FROM stop;

   -- SELECT stop_id,stop_name,st_x(st_aste'index de la route du cheminxt(stop_coord)), st_y(st_astext(stop_coord)) FROM stop  WHERE ST_DWithin(stop_coord, ST_GeographyFromText('SRID=4326;POINT(48.11187 -1.68481)'), 200);

-- affichage de la route
/*select * from route where route_id in (
	-- determine les routes possibles en focntion des trip et des services disponible (ici le lundi)
	select route_id from trip where trip_id in  (
		-- détermine les trip possibles en fonction de l'heure  et de l'arret
		select trip_id from stop_time  where stop_id in ( 
			-- determine le point de départ
			SELECT stop_id FROM stop  WHERE ST_DWithin(stop_coord, ST_GeographyFromText('SRID=4326;POINT(48.11187 -1.68481)'), 200) 
		)
		and arrival_time::time > '14:05' and arrival_time::time < '14:09:00'
	) and  service_id in (3,16,17,18,19,20,21)
);*/


select stt.trip_id,  trip_headsign, departure_time, stop_name, stop_desc, stop_coord, stop_sequence from stop_time stt join stop st on stt.stop_id = st.stop_id join trip t on t.trip_id = stt.trip_id   where stt.trip_id in (
	-- determine les routes possibles en focntion des trip et des services disponible (ici le lundi)
	select trip_id from trip where trip_id in  (
		-- détermine les trip possibles en fonction de l'heure  et de l'arret
		select trip_id from stop_time  where stop_id in ( 
			-- determine le point de départ
			SELECT stop_id FROM stop  WHERE ST_DWithin(stop_coord, ST_GeographyFromText('SRID=4326;POINT(48.109970 -1.679217)'), 200) 
		)
		and arrival_time::time > '14:05' and arrival_time::time < '15:09:00'
	) and  service_id in (3,16,17,18,19,20,21)
)  order by trip_id,stop_sequence::int  ;






