package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"

	svg "github.com/marcelmue/svgo"

	"github.com/n3wscott/kremanaughts/pkg/crema"
	gg "github.com/n3wscott/kremanaughts/pkg/graph"
	"github.com/n3wscott/kremanaughts/pkg/path"
)

func main() {
	http.Handle("/graph", http.HandlerFunc(graph))
	http.Handle("/", http.HandlerFunc(circle))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var HA = math.Pi * 2.0 / 7.0

func circle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	s := svg.New(w)
	s.Start(500, 500)

	//s.Circle(250, 250, 100, "fill:none;stroke:black")

	c := crema.New(100.0, 2.0, 5.0)

	s.Circle(250, 250, int(c.R), "fill:none;stroke:black")
	s.Circle(250, 250, int(c.R)+20, "fill:none;stroke:black")
	//s.Circle(250-int(c.Block.X), 250-int(c.Block.Y), int(c.Block.R), "fill:none;stroke:red")
	//s.Path(Heptagon(250-c.Block.X, 250.0-c.Block.Y, c.Block.R, rand.Float64()*math.Pi), "fill:none;stroke:blue")

	for _, a := range c.Points {
		//s.Circle(250-int(a.X), 250-int(a.Y), int(a.R), "fill:none;stroke:blue")
		s.Path(Heptagon(250-a.X, 250.0-a.Y, a.R, rand.Float64()*math.Pi), "fill:none;stroke:brown")
	}

	//dx := 250.0
	//dy := 250.0
	//r := 80.0
	//for i := 0.0; i < 20.0; i += 1.0 {
	//	a := (math.Pi * 2.0 / 360.0) * i
	//	x := r*math.Sin(a) + dx
	//	y := r*math.Cos(a) + dy
	//	s.Path(Heptagon(x, y, 10, 0.0), "fill:none;stroke:black;stroke-width:3")
	//}

	//s.Path(Heptagon(250.0, 230.0, 10.0, 0.0), "fill:none;stroke:black;stroke-width:3")
	//s.Path(Heptagon(250.0, 250.0, 10.0, HA/2), "fill:none;stroke:black;stroke-width:3")

	s.End()
}

// d (straight line on a circle) = 2 * r * sin(angle/2)

func Heptagon(x, y, r, rotate float64) string {
	p := path.New()
	for i := 0.0; i < 7.0; i += 1.0 {
		a := i*HA + HA/2.0 + rotate
		x := r*math.Sin(a) + x
		y := r*math.Cos(a) + y
		if i == 0.0 {
			p.Start(x, y)
		} else {
			p.LineABS(x, y)
		}
	}
	p.Connect()

	return p.String()
}

// 51.42857142857 angle at center from two points.

func graph(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	s := svg.New(w)
	s.Start(1200, 800)

	s.Circle(200, 300, 10, "fill:none;stroke:black")
	s.Circle(1000, 300, 10, "fill:none;stroke:black")

	s.Circle(400, 50, 10, "fill:none;stroke:red")
	s.Circle(600, 300, 10, "fill:none;stroke:red")

	var v1 *gg.Vector
	var v2 *gg.Vector
	var v3 *gg.Vector
	var v4 *gg.Vector

	s1 := gg.NewBox(250, 250, 100, 300)
	s.Rect(250, 250, 100, 300, "fill:none;stroke:blue")
	for d, edge := range s1.Points() {
		if d == 0 {
			xx := edge
			v1 = &xx
		}
		if d == 3 {
			xx := edge
			v3 = &xx
		}
		s.Circle(int(edge.X), int(edge.Y), 10, "fill:none;stroke:orange")
		fmt.Printf("%+v\n", v1)
	}

	s2 := gg.NewBox(600, 600, 300, 100)
	s.Rect(600, 600, 300, 100, "fill:none;stroke:blue")
	for d, edge := range s2.Points() {
		if d == 2 {
			xx := edge
			v2 = &xx
		}
		if d == 3 {
			xx := edge
			v4 = &xx
		}
		s.Circle(int(edge.X), int(edge.Y), 10, "fill:none;stroke:orange")
	}

	if v1 != nil && v2 != nil {
		p := path.New()
		p.Start(v1.X, v1.Y)

		dx1, dy1 := v1.Head(100)
		dx2, dy2 := v2.Head(100)

		hx := v1.X + (v2.X-v1.X)/2.0
		hy := v1.Y + (v2.Y-v1.Y)/2.0

		dhx := hx
		dhy := hy - 50

		p.CurveABS(dx1, dy1, dhx, dhy, hx, hy)
		p.SymmetricABS(dx2, dy2, v2.X, v2.Y)

		s.Path(p.String(), "fill:none;stroke:blue")
	}

	if v3 != nil && v4 != nil {
		p := path.New()
		p.Start(v3.X, v3.Y)

		dx1, dy1 := v3.Head(100)
		dx2, dy2 := v4.Head(100)

		hx := v3.X + (v4.X-v3.X)/2.0
		hy := v3.Y + (v4.Y-v3.Y)/2.0

		dhx := hx
		dhy := hy - 200

		p.CurveABS(dx1, dy1, dhx, dhy, hx, hy)
		p.SymmetricABS(dx2, dy2, v4.X, v4.Y)

		s.Path(p.String(), "fill:none;stroke:blue")
	}

	{
		p := path.New()

		// M200,300 Q400,50 600,300 T1000,300

		p.Start(200, 300)
		p.QuadraticABS(400, 50, 600, 300)
		p.QuadraticSmoothABS(1000, 300)

		//p.SymmetricABS(400, 250, 400, 400)
		//p.QuadraticABS(400, 250, 400, 400)
		s.Path(p.String(), "fill:none;stroke:brown")
	}
	s.End()
}
