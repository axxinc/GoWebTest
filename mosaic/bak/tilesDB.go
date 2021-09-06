package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
)

func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db...")

	db := make(map[string][3]float64)

	fileDir, _ := ioutil.ReadDir("tiles")

	for _, f := range fileDir {

		name := "tiles/" + f.Name()
		file, err := os.Open(name)
		fmt.Println(name)
		if err == nil {
			img, _, err := image.Decode(file)

			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB - ", name, err)
			}
		} else {
			fmt.Println("cannot open file - ", name, err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	return db
}
