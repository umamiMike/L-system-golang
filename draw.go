package main

import (
	"github.com/fogleman/gg"
)

/*
what i want to draw is based on the grammer produced above
a will move forward

*/

type Vector struct {
	X, Y, Z float64
}

type Color struct {
	R, G, B, A float64
}

type Canvas struct {
	w  int
	h  int
	dc *gg.Context
}

func newCanvas(w int, h int) Canvas {

	return Canvas{
		w: w,
		h: h,

		dc: gg.NewContext(w, h),
	}

}

func drawImage(fname string, w int, h int, instructions string ) string {
	iter := 7000
	canvas := newCanvas(w, h)
	dc := canvas.dc
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	for i := 0; i < iter; i++ {
		////create random color

		//r := rand.Float64()
		////g := rand.Float64()
		////b := rand.Float64()
		//a := rand.Float64()*0.5 + 0.5
		//dc.SetRGBA(r, r, r, a)
		//w := rand.Float64()*6 + 1
		//dc.SetLineWidth(w)
		////create random position

		//x1 := rand.Float64() * float64(canvas.w)
		//y1 := rand.Float64() * float64(canvas.h)

		//x2 := rand.Float64() * float64(canvas.w)
		//y2 := rand.Float64() * float64(canvas.h)

		//dc.DrawLine(x1, y1, x2, y2)

		//dc.Stroke()
	}
	dc.SavePNG(fname)
	return fname

}

func drawSpike(fname string, w int, h int, instructions string ) string {
	iter := 7000
	canvas := newCanvas(w, h)
	dc := canvas.dc
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	for i := 0; i < iter; i++ {
		////create random color

		//r := rand.Float64()
		////g := rand.Float64()
		////b := rand.Float64()
		//a := rand.Float64()*0.5 + 0.5
		//dc.SetRGBA(r, r, r, a)
		//w := rand.Float64()*6 + 1
		//dc.SetLineWidth(w)
		////create random position

		//x1 := rand.Float64() * float64(canvas.w)
		//y1 := rand.Float64() * float64(canvas.h)

		//x2 := rand.Float64() * float64(canvas.w)
		//y2 := rand.Float64() * float64(canvas.h)

		//dc.DrawLine(x1, y1, x2, y2)

		//dc.Stroke()
	}
	dc.SavePNG(fname)
	return fname

}
