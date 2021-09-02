package mosaic

import (
	"image"
	"image/color"
)

func resize(img image.Image, newWidth int) image.NRGBA {
	bounds := img.Bounds()

	// dx 宽度
	ratio := bounds.Dx() / newWidth

	// 返回一个矩形Rectangle{Pt(x0, y0), Pt(x1, y1)}。
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.Y/ratio, bounds.Max.X/ratio, bounds.Max.Y/ratio))

	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Min.X; x, i = x+ratio, i+1 {
			r, g, b, a := img.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}
