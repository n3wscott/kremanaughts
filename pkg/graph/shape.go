package graph

import "math"

type Node interface {
	Points() []Vector
	Center() (float64, float64)
}

type Point struct {
	X float64
	Y float64
}

type Vector struct {
	Point
	D float64 // direction, radians
}

func (v *Vector) Head(mag float64) (float64, float64) {
	x := v.X + mag*math.Cos(v.D)
	y := v.Y + mag*math.Sin(v.D)
	return x, y
}

type Box struct {
	center Point
	width  float64
	height float64
}

func (b Box) Points() []Vector {
	v := make([]Vector, 0)
	w := b.width / 2.0
	h := b.height / 2.0
	for _, d := range []float64{0, math.Pi / 2, math.Pi, 3 * math.Pi / 2} {
		v = append(v, Vector{
			Point: Point{
				X: b.center.X + w*math.Cos(d),
				Y: b.center.Y + h*math.Sin(d),
			},
			D: d,
		})
	}
	return v
}

func (b Box) Center() (float64, float64) {
	return b.center.X, b.center.Y
}

func NewBox(x, y, width, height float64) Node {
	return &Box{
		center: Point{
			X: x + width/2.0,
			Y: y + height/2.0,
		},
		width:  width,
		height: height,
	}
}
