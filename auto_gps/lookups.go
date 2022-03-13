package auto_gps

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

	sql := `select city, state, tz, 
		Distance(GeomFromText('POINT(?, ?)', 4326), location) as distance,
		Azimuth(GeomFromText('POINT(?, ?)', 4326), location) as direction
		from citylocations
		order by distance asc limit 3`

	result, err := Agps.Spatial.Query(sql, lat, lon, lat, lon)

	if err != nil {
		log.Printf("Get Locations: %+v", err)
	}

	curresult := 0

	for result.Next() {
		var location Location = Location{}
		result.Scan(&location)

		if curresult == 0 {
			Agps.Tz = location.Tz
			curresult = 1
		}

		log.Printf("Row: %+v", location)
		cities = append(cities, fmt.Sprintf("%s, %s %f", location.City, location.State, location.Distance))
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
		"DROP TABLE IF EXISTS citylocations",
		"CREATE TABLE citylocations (id INTEGER PRIMARY KEY AUTOINCREMENT, city text, state text, tz text);",
		"SELECT AddGeometryColumn('citylocations', 'location', 4326, 'POINT', 'XY', 0);",
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

	sql := `insert into citylocations (city, state, tz, location) values ( ?, ?, ?, GeomFromText('POINT( ? ? )', 4326));`

	tx, _ := Agps.Spatial.Begin()

	sth, _ := tx.Prepare(sql)

	for i, line := range data {
		if i%1000 == 0 {
			log.Printf("%0d lines\n", i)
		}
		sth.Exec(sql, line[0], line[3], line[13], line[6], line[7])
	}

	tx.Commit()

	f.Close()
}
