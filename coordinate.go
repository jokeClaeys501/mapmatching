package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type coordinate struct {
	longitude float64
	latitude float64
}

func getCoordinates() {
		var coordinates []coordinate

	f, err := excelize.OpenFile("longlat.xlsx")
	if err != nil {
        	log.Fatal(err)
    	}

	for i := 0; i < 3; i++ {
		long, err := f.GetCellValue("Sheet1", fmt.Sprint("A", i+2))
		lat, err := f.GetCellValue("Sheet1", fmt.Sprint("B", i+2))
		if err != nil {
        	log.Fatal(err)
    	}
		longFloat, err := strconv.ParseFloat(long, 64)
		latFloat, err := strconv.ParseFloat(lat, 64)
		coordinateTemp := coordinate{
			longitude: longFloat,
			latitude: latFloat,
		}
		coordinates = append(coordinates, coordinateTemp)
	}

    
	fmt.Println(coordinates)
}