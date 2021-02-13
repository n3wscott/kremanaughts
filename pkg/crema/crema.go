package crema

import (
	"math"
	"math/rand"
	"time"
)

func New(r, min, max float64) *Crema {
	c := &Crema{
		R:      r,
		Points: make([]Circle, 0),
	}
	c.Generate(min, max)
	return c
}

func (c *Crema) Generate(min, max float64) {
	c.Block = c.RandCircle(c.R/3.0, c.R/1.2)

	t := 0
	for i := 0; i < 1000; {
		t++
		a := c.RandCircle(min, max)
		if c.Validate(a) {
			c.Points = append(c.Points, a)
			i++
		}
		if t > 100000 {
			break
		}
	}
}

func (c *Crema) Validate(a Circle) bool {
	if overlaps(c.Block, a) {
		return false
	}
	for _, b := range c.Points {
		if overlaps(a, b) {
			return false
		}
	}
	return true
}

func (c *Crema) RandCircle(minR, maxR float64) Circle {
	x := Circle{
		R: random(minR, maxR),
	}

	angle := random(0, math.Pi*2.0)
	shift := random(0, c.R-x.R)

	x.X = shift * math.Sin(angle)
	x.Y = shift * math.Cos(angle)

	return x
}

func overlaps(a, b Circle) bool {
	x := math.Abs(a.X - b.X)
	y := math.Abs(a.Y - b.Y)
	h := math.Sqrt(x*x + y*y)
	if h < a.R+b.R {
		return true
	}
	return false
}

func random(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Crema struct {
	R      float64
	Block  Circle
	Points []Circle
}

type Circle struct {
	X float64
	Y float64
	R float64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
