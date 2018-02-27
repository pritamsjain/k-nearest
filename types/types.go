package types

import "math"

type RecordReader interface {
	New(connInfo string)
	GetAllRecords()[]Point
}
// Earth radius in kilometers.
const radius = 6371

type Point interface {
	GetId() uint64
	// to support more dimentions
	Dim() int
	// Return the value X_{dim}, dim is started from 0
	GetValue(dim int) float64
	// Return the distance between two points
	Distance(p Point) float64
	// Return the distance between the point and the plane X_{dim}=val
	PlaneDistance(val float64, dim int) float64
}
//impl of above Point interface
type GeoPoints struct {
	Id          uint64
	Coordinates []float64
}
func  (p *GeoPoints) GetId() uint64{
	return p.Id
}

func (p *GeoPoints) Dim() int {
	return len(p.Coordinates)
}

func (p *GeoPoints) GetValue(dim int) float64 {
	return p.Coordinates[dim]
}
//distance calculation https://en.wikipedia.org/wiki/Great-circle_distance (spherical law of cosines)
func (p *GeoPoints) Distance(other Point) float64 {
		s1, c1 := math.Sincos(rad(p.GetValue(0)))
		s2, c2 := math.Sincos(rad(other.GetValue(0)))
		clong := math.Cos(rad(p.GetValue(1) - other.GetValue(1)))
		return radius * math.Acos(s1*s2+c1*c2*clong)
}

func (p *GeoPoints) PlaneDistance(val float64, dim int) float64 {
	tmp := p.GetValue(dim) - val
	return tmp * tmp
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}