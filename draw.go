package main

import (
	"github.com/fogleman/gg"
	"math/rand"
)

/*
what i want to draw is based on the grammer produced above
a will move forward

*/
var canvas = struct {
	W int
	H int
}{
	W: 1024,
	H: 1024,
}

func draw(s string) {

	dc := gg.NewContext(canvas.W, canvas.H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	for i := 0; i < 1000; i++ {
		x1 := rand.Float64() * canvas.W
		y1 := rand.Float64() * canvas.H
		x2 := rand.Float64() * canvas.W
		y2 := rand.Float64() * canvas.H
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		w := rand.Float64()*4 + 1
		dc.SetRGBA(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
	dc.SavePNG(s)
}
