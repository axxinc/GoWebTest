package mosaic

import "image"

func averageColor(img image.Image) [3]float64 {
	// 返回图像范围
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0

	// At(Bounds().Min.X, Bounds().Min.Y)返回网格左上角像素的色彩
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) 返回网格右下角像素的色彩
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.Y; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}
