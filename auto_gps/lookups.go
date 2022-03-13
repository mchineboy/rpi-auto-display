package auto_gps

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Location struct {
	City      string
	State     string
	Tz        string
	Distance  float64
	Direction float64
}

func (Agps *AutoGps) FindNearestTowns(lat float64, lon float64) []string {
	var cities []string
	log.Printf("%0.3f, %0.3f", lon, lat)
	sql := fmt.Sprintf(`select City, State, Tz, 
		Distance(GeomFromText('POINT(%f %f)', 4326), location, 0) as Distance,
		Azimuth(GeomFromText('POINT(%f %f)', 4326), location) as Direction
		from citylocations
		order by distance asc limit 3`, lon, lat, lon, lat)

	result, err := Agps.Spatial.Query(sql)

	if err != nil {
		log.Printf("Get Locations: %+v", err)
	}

	curresult := 0

	for result.Next() {
		var city string
		var state string
		var tz string
		var distance float64
		var direction float64
		result.Scan(&city, &state, &tz, &distance, &direction)

		direction = direction * 180 / math.Pi
		compass := ""
		switch {
		case direction <= 22.5:
			compass = "N"
		case direction < 67.5:
			compass = "NE"
		case direction <= 112.5:
			compass = "E"
		case direction < 157.5:
			compass = "SE"
		case direction <= 202.5:
			compass = "S"
		case direction < 247.5:
			compass = "SW"
		case direction <= 292.5:
			compass = "W"
		case direction < 337.5:
			compass = "NW"
		default:
			compass = "N"
		}
		if curresult == 0 {
			Agps.Tz = tz
			curresult = 1
		}

		cities = append(cities, fmt.Sprintf("%s, %s %0.1f %s", city, state,
			distance/1609, compass))
	}
	log.Printf("%+v", cities)
	return cities
}

func (Agps *AutoGps) BuildDatabase() {
	f, err := os.Open("/data/uscities.csv")
	if err != nil {
		log.Printf("Open US Cities File: %+v", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Printf("CSV Reader: %+v\n", err)
	}
	q := "SELECT InitSpatialMetaData(1);"
	Agps.Spatial.Exec(q)

	sqls := []string{
		"SELECT DropGeoTable('citylocations')",
		"CREATE TABLE citylocations (id INTEGER PRIMARY KEY AUTOINCREMENT, city text, state text, tz text);",
		"SELECT AddGeometryColumn('citylocations', 'location', 4326, 'POINT', 'XY', 1);",
		"SELECT CreateSpatialIndex('citylocations', 'location');",
	}
	for _, sql := range sqls {
		log.Println(sql)
		result, err := Agps.Spatial.Exec(sql)
		if err != nil {
			log.Printf("Error executing %s : %+v", sql, err)
		}
		log.Printf("%+v", result)
	}

	ctx := context.Background()

	tx, _ := Agps.Spatial.BeginTx(ctx, nil)

	if err != nil {
		log.Printf("Prepare: %+v", err)
	}
	for i, line := range data {
		if i == 0 {
			continue
		}
		if i%1000 == 0 && i > 0 {
			log.Printf("%0d lines..\n", i)
		}

		lon, _ := strconv.ParseFloat(line[7], 64)
		lat, _ := strconv.ParseFloat(line[6], 64)

		sql := fmt.Sprintf(
			`insert into citylocations (city, state, tz, location) values ( ?, ?, ?, GeomFromText('POINT(%f %f)', 4326));`,
			lon, lat)
		_, err := tx.ExecContext(ctx, sql, line[0], line[2], line[13])

		if err != nil {
			log.Panicf("Error on Insert: %+v %+v", err, sql)
		}

	}

	tx.Commit()

	result, err := Agps.Spatial.Exec(`vacuum; analyze;`)
	if err != nil {
		log.Printf("Vacuum Analyze: %+v", err)
	}

	log.Println(result.RowsAffected())

	f.Close()
}
