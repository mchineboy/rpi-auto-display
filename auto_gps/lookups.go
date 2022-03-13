package auto_gps

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
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
	sql := fmt.Sprintf(`select city, state, tz, 
		Distance(GeomFromText('POINT(%f %f)', 4326), location, 1) as distance,
		Azimuth(GeomFromText('POINT(%f %f)', 4326), location) as direction
		from citylocations
		order by distance asc limit 3`, lon, lat, lon, lat)

	result, err := Agps.Spatial.Query(sql)

	if err != nil {
		log.Printf("Get Locations: %+v", err)
	}

	curresult := 0

	for result.Next() {
		location := map[string]interface{}{}
		result.Scan(&location)

		if curresult == 0 {
			Agps.Tz = fmt.Sprintf("%s", location["tz"])
			curresult = 1
		}

		log.Printf("Row: %s %s %f %f",
			location["city"], location["state"],
			location["distance"], location["direction"])
		cities = append(cities, fmt.Sprintf("%s, %s %f", location["city"], location["state"],
			location["distance"]))
	}

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
		_, err := tx.ExecContext(ctx, sql, line[0], line[3], line[13])

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
