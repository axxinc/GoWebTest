package main

import "math"

func nearest(target [3]float64, db *map[string][3]float64) string {
	var fileName string
	smallest := 1000000.0

	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			fileName, smallest = k, dist
		}
	}
	delete(*db, fileName)
	return fileName
}

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(Sq(p2[0]-p1[0]) + Sq(p2[1]-p1[1]+Sq(p2[2]-p1[1])))
}

func Sq(n float64) float64 {
	return n * n
}
