package rb

import (
	"image"
	"image/color"
	"image/draw"
)

func Circle(p image.Point, r int, red bool, img draw.Image, col color.Color) {
	for m := p.X - r; m <= p.X+r; m++ {
		for n := p.Y - r; n <= p.Y+r; n++ {
			xx, yy, rr := float64(m-p.X)+0.5, float64(n-p.Y)+0.5, float64(r)
			if xx*xx+yy*yy < rr*rr {
				img.Set(m, n, col)
			}
		}
	}
}
