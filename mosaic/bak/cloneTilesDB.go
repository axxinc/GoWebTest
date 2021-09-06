package main

var TILESDB map[string][3]float64

func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)

	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}
