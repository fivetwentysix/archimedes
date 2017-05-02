package main

import (
	"encoding/csv"
	"log"

	"fmt"
	"io"
	"os"

	forecast "github.com/mlbright/forecast/v2"
)

func getWeather(key, zipcode string) (summary string, temperature float64) {
	lat, lon := getCoords(zipcode)

	f, err := forecast.Get(key, lat, lon, "now", forecast.US, forecast.English)

	if err != nil {
		log.Fatal(err)
	}
	return f.Currently.Summary, f.Currently.Temperature
}

func getCoords(zip string) (lat, lon string) {
	file, err := os.Open("data/zip.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	lineCount := 0
	for {

		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if record[0] == zip {
			latitude := record[1]
			longitude := record[2]
			return latitude, longitude
		}
		lineCount++
	}

	// Some valid value if no match is found.
	return "0", "0"
}
