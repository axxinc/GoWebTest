package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image")
	defer file.Close()

	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	img, _, _ := image.Decode(file)
	bounds := img.Bounds()

	newImg := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))

	db := cloneTilesDB()

	sp := image.Point{0, 0}

	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {

			r, g, b, _ := img.At(x, y).RGBA()
			clolr := [3]float64{float64(r), float64(g), float64(b)}
			// 从db中找到对应颜色的照片名称
			nearest := nearest(clolr, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img1, _, err := image.Decode(file)

				if err == nil {
					t := resize(img1, tileSize)
					tile := t.SubImage(t.Bounds())

					// 新图像像素大小
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)

					draw.Draw(newImg, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("mian.go 60 err:", err, nearest)
				}

			} else {

				fmt.Println("main.go 65 file err:", err, nearest)
			}
			file.Close()
		}
	}

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, img, nil)
	imgstr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newImg, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())

	ti := time.Now()

	images := map[string]string{
		"original": imgstr,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v", ti.Sub(t0)),
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)

	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}

	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")

	server.ListenAndServe()
}
