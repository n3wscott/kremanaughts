package main

import (
	"github.com/n3wscott/kremanaughts/pkg/crema"
	"github.com/n3wscott/kremanaughts/pkg/path"
	"log"
	"math"
	"math/rand"
	"net/http"

	svg "github.com/marcelmue/svgo"
)

func main() {
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
