package main

import (
	"flag"
	"fmt"
	"k-nearest/csv"
	"k-nearest/lib"
	"k-nearest/types"
	"strings"
)

var source = flag.String("source", "geoData.csv", "can be path to csv file or url to connect db")

func main() {
	flag.Parse()
	fmt.Println("Recieved source", *source)
	var recordReader types.RecordReader
	//as per source we can decide reader ie csv reader or database reader
	if strings.HasSuffix(*source, ".csv") {
		recordReader = &csv.CsvReader{*source}
	}
	allInputpts := recordReader.GetAllRecords()
	fmt.Println("Read all input points, Total :",len(allInputpts) )

	fmt.Println("Building tree" )
	tree := kdtree.New(allInputpts)

	officeLocation := types.GeoPoints{}
	officeLocation.Coordinates = []float64{51.925146, 4.478617}

	fmt.Println("5 nearest locations are:" )
	neighbours := tree.KNearestNeibors(&officeLocation, 5)
	printLocationWithDistance(neighbours,&officeLocation)

	fmt.Println("5 Further Most locations are:" )
	logDistanceNeigbours := tree.KFurthermostNeibors(&officeLocation, 5)
	printLocationWithDistance(logDistanceNeigbours,&officeLocation)
}

func printLocationWithDistance(points []types.Point,from types.Point)  {
	for _, p := range points {
		fmt.Printf("Point id %d: (%f,%f) distance %v km\n",
			p.GetId(), p.GetValue(0),p.GetValue(1),p.Distance(from))
	}
}