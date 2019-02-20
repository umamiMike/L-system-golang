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
	W float64
	H float64
}{
	W: 1024,
	H: 1024,
}

func drawSpike(s string, iter int) string {

	dc := gg.NewContext(int(canvas.W), int(canvas.H))
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	for i := 0; i < iter; i++ {
		//create random color
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		dc.SetRGBA(r, g, b, a)
		w := rand.Float64()*4 + 1
		dc.SetLineWidth(w)
		//create random position
		x1 := rand.Float64() * canvas.W
		x2 := rand.Float64() * canvas.W
		y1 := rand.Float64() * canvas.H
		y2 := rand.Float64() * canvas.H
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
	dc.SavePNG(s)
	return s
}
