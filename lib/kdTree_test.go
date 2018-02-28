package kdtree

import (
	"testing"
	"k-nearest/types"
)


func TestKNearestNeibhors(t *testing.T) {
	// case 1
	{
		p1 := types.GeoPoints{1,[]float64{0.0,0.0}}
		p2 := types.GeoPoints{1,[]float64{0.0,1.0}}
		p3 := types.GeoPoints{1,[]float64{0.0,2.0}}
		p4 := types.GeoPoints{1,[]float64{0.0,3.0}}
		points := make([]types.Point, 0)
		points = append(points, &p1)
		points = append(points, &p2)
		points = append(points, &p3)
		points = append(points, &p4)
		tree := New(points)
		currentloc:=types.GeoPoints{1,[]float64{-1.0,-1.0}}
		ans := tree.KNearestNeibors(&currentloc, 2)
		compareResult(t, ans, &p1, &p2)
	}
}

func TestKfurthermostNeibhors(t *testing.T) {
	{
		p1 := types.GeoPoints{1,[]float64{0.0,0.0}}
		p2 := types.GeoPoints{1,[]float64{0.0,1.0}}
		p3 := types.GeoPoints{1,[]float64{0.0,2.0}}
		p4 := types.GeoPoints{1,[]float64{0.0,3.0}}
		points := make([]types.Point, 0)
		points = append(points, &p1)
		points = append(points, &p2)
		points = append(points, &p3)
		points = append(points, &p4)
		tree := New(points)
		currentloc:=types.GeoPoints{1,[]float64{-1.0,-1.0}}
		ans := tree.KFurthermostNeibors(&currentloc, 2)
		compareResult(t, ans, &p4, &p3)
	}
}
func equal(p1 types.Point, p2 types.Point) bool {
	for i := 0; i < p1.Dim(); i++ {
		if p1.GetValue(i) != p2.GetValue(i) {
			return false
		}
	}
	return true
}

func compareResult(t *testing.T, ans []types.Point, points ...types.Point) {
	if len(ans) != len(points) {
		t.Fatal("Nearest Neibhors result length error",len(ans), len(points))
	}
	for i := 0; i < len(ans); i++ {
		if !equal(ans[i], points[i]) {
			t.Error("Nearest Neibhors results are wrong")
		}
	}
}
