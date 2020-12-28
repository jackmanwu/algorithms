package rb

import (
	"image/color"
	"image/draw"
	"math"
)

func HLine(x1, y, x2 int, img draw.Image, col color.Color) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

func VLine(x, y1, y2 int, img draw.Image, col color.Color) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

func Line(x1, y1, x2, y2 int, img draw.Image, col color.Color) {
	dx := math.Abs(float64(x2 - x1))
	var sx int
	if x1 < x2 {
		sx = 1
	} else {
		sx = -1
	}
	dy := -math.Abs(float64(y2 - y1))
	var sy int
	if y1 < y2 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx + dy
	for {
		img.Set(x1, y1, col)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 >= dy {
			err += dy
			x1 += sx
		}
		if e2 <= dx {
			err += dx
			y1 += sy
		}
	}
}

func LineLow(x1, y1, x2, y2 int, img draw.Image, col color.Color) {
	dx := x2 - x1
	dy := y2 - y1
	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	d := 2*dy - dx
	y := y1

	for x := x1; x <= x2; x++ {
		img.Set(x, y, col)
		if d > 0 {
			y += yi
			d += 2 * (dy - dx)
		} else {
			d += 2 * dy
		}
	}
}

func LineHigh(x1, y1, x2, y2 int, img draw.Image, col color.Color) {
	dx := x2 - x1
	dy := y2 - y1
	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	d := 2*dx - dy
	x := x1
	for y := y1; y <= y2; y++ {
		img.Set(x, y, col)
		if d > 0 {
			x += xi
			d += 2 * (dx - dy)
		} else {
			d += 2 * dx
		}
	}
}

func LineV2(x1, y1, x2, y2 int, img draw.Image, col color.Color) {
	if math.Abs(float64(y2-y1)) < math.Abs(float64(x2-x1)) {
		if x1 > x2 {
			LineLow(x2, y2, x1, y1, img, col)
		} else {
			LineLow(x1, y1, x2, y2, img, col)
		}
	} else {
		if y1 > y2 {
			LineHigh(x2, y2, x1, y1, img, col)
		} else {
			LineHigh(x1, y1, x2, y2, img, col)
		}
	}
}
