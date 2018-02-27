package csv

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"k-nearest/types"
	"encoding/csv"
)
//implementing type/RecordReader
type CsvReader struct {
	CsvFilePath string
}

func (c *CsvReader)New(connInfo string)   {
	c.CsvFilePath = connInfo
}

func (c *CsvReader) GetAllRecords()[]types.Point  {

	f, err := os.Open(c.CsvFilePath)

	if err!=nil{
		fmt.Errorf("error while opening file %v",err)
		return nil
	}

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	records, err := r.ReadAll()
	if err!=nil{
		fmt.Errorf("error while reading file %v",err)
		return nil
	}
	points := make([]types.Point,0)
	for i:=1;i<len(records);i++ {
		points = append(points, preparePoint(records[i]))

	}
	return points
}

func preparePoint(recordLine []string)(types.Point)  {
	var pt = types.GeoPoints{}
	id,err:= strconv.ParseUint(recordLine[0],10,64)
	if err!=nil{
		fmt.Errorf("error while interpreting data %v, got %v",err,recordLine[0])

	}
	lat,err:= strconv.ParseFloat(recordLine[1],64)
	if err!=nil{
		fmt.Errorf("error while interpreting data %v, got %v",err,recordLine[1])
	}
	long,err:= strconv.ParseFloat(recordLine[2],64)
	if err!=nil{
		fmt.Errorf("error while interpreting data %v, got %v",err,recordLine[2])
	}
	pt.Id = id
	pt.Coordinates = []float64{lat,long}
	return &pt
}
