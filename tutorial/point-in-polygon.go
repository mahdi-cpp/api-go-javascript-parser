package tutorial

import "fmt"

type Point struct {
	X float64
	Y float64
}

func isPointInPolygon(point Point, polygon []Point) bool {
	oddNodes := false
	j := len(polygon) - 1

	for i := 0; i < len(polygon); i++ {
		if (polygon[i].Y < point.Y && polygon[j].Y >= point.Y) || (polygon[j].Y < point.Y && polygon[i].Y >= point.Y) {
			if polygon[i].X+(point.Y-polygon[i].Y)/(polygon[j].Y-polygon[i].Y)*(polygon[j].X-polygon[i].X) < point.X {
				oddNodes = !oddNodes
			}
		}
		j = i
	}

	return oddNodes
}

func MainPointInPolygon() {
	// Define a polygon (list of points)
	polygon := []Point{
		{1, 1},
		{4, 8},
		{3, 3},
		{3, 1},
	}

	// Define a point to check
	point := Point{2, 3}

	// Check if the point is inside the polygon
	if isPointInPolygon(point, polygon) {
		fmt.Println("Point is inside the polygon")
	} else {
		fmt.Println("Point is outside the polygon")
	}
}
