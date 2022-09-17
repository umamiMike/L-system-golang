package main

import (
	"fmt"
	"math/rand"

	"github.com/fogleman/gg"
)

/*
what i want to draw is based on the grammer produced above
a will move forward


Binary tree drawing rules
    0: draw a line segment ending in a leaf
    1: draw a line segment
    [: push position and angle, turn left 45 degrees
    ]: pop position and angle, turn right 45 degrees

*/

func move_forward() {
	fmt.Println("moving forward")

}

func draw_forward() {
	fmt.Println("drawing forward")

}

func turn_left() {
	fmt.Println("turning left")

}
func turn_right() {
	fmt.Println("turning right")

}

var pencil = map[string]func(){"a": move_forward, "b": draw_forward, "[": turn_left, "]": turn_right}

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

func drawImage(fname string, w int, h int, instructions string) string {
	canvas := newCanvas(w, h)
	dc := canvas.dc
	// dc.SetRGB(0, 0, 0)
	// dc.Clear()
	for _, v := range instructions {
		//x2 := rand.Float64() * float64(canvas.w)
		//y2 := rand.Float64() * float64(canvas.h)
		rw := rand.Float64() * float64(canvas.w)
		rh := rand.Float64() * float64(canvas.h)
		dc.SetRGBA(1, 1, 1, .3)
		dc.LoadFontFace("/usr/share/fonts/truetype/freefont/FreeMono.ttf", 22)
		dc.DrawString(string(v), rw, rh)

		for _, v := range instructions {
			pencil[string(v)]()

		}
		//using the pencil
		// translate the grammer using the functions

	}
	dc.SavePNG(fname)
	return fname

}

//draw a line
