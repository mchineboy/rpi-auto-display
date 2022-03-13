package auto_gps

import (
	"context"
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
	log.Printf("%0.3f, %0.3f", lon, lat)
	sql := `select city, state, tz, 
		Distance(MakePoint(?, ?), location, 1) as distance,
		Azimuth(MakePoint(?, ?), location) as direction
		from citylocations
		order by distance asc limit 3`

	result, err := Agps.Spatial.Query(sql, lon, lat, lon, lat)

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

	sql := `insert into citylocations (city, state, tz, location) values ( ?, ?, ?, MakePoint(?, ?));`

	ctx := context.Background()

	tx, _ := Agps.Spatial.BeginTx(ctx, nil)

	sth, err := tx.PrepareContext(ctx, sql)

	if err != nil {
		log.Printf("Prepare: %+v", err)
	}
	defer sth.Close()
	for i, line := range data {
		if i%1000 == 0 {
			log.Printf("%0d lines.. commit\n", i)
			tx.Commit()
		}
		_, err := sth.ExecContext(ctx, sql, line[0], line[3], line[13], line[7], line[6])
		if err != nil {
			log.Panicf("Error on Insert: %+v", err)
		}

	}
	sth.Close()
	tx.Commit()

	result, err := Agps.Spatial.Exec(`vacuum; analyze;`)
	if err != nil {
		log.Printf("Vacuum Analyze: %+v", err)
	}

	log.Println(result.RowsAffected())

	f.Close()
}
