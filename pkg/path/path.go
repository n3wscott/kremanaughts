package path

import (
	"fmt"
	"strings"
)

// Comments from https://css-tricks.com/svg-path-syntax-illustrated-guide/
// Better docs here: https://www.w3.org/TR/SVG/paths.html#PathDataCubicBezierCommands

func New() *Path {
	return &Path{b: strings.Builder{}}
}

type Path struct {
	b strings.Builder
}

func (p *Path) String() string {
	return p.b.String()
}

/*

Straight Lines

M x,y	Move to the absolute coordinates x,y
m x,y	Move to the right x and down y (or left and up if negative values)
L x,y	Draw a straight line to the absolute coordinates x,y
l x,y	Draw a straight line to a point that is relatively right x and down y (or left and up if negative values)
H x		Draw a line horizontally to the exact coordinate x
h x		Draw a line horizontally relatively to the right x (or to the left if a negative value)
V y		Draw a line vertically to the exact coordinate y
v y		Draw a line vertically relatively down y (or up if a negative value)
Z		Draw a straight line back to the start of the path
*/

// M x,y	Move to the absolute coordinates x,y
func (p *Path) Start(x, y float64) *Path {
	p.b.WriteString(fmt.Sprintf("M %.4f, %.4f ", x, y))
	return p
}

// m x,y	Move to the right x and down y (or left and up if negative values)
func (p *Path) Move(x, y float64) *Path {
	p.b.WriteString(fmt.Sprintf("m %.4f, %.4f ", x, y))
	return p
}

// L x,y	Draw a straight line to the absolute coordinates x,y
func (p *Path) LineABS(x, y float64) *Path {
	p.b.WriteString(fmt.Sprintf("L %.4f, %.4f ", x, y))
	return p
}

// l x,y	Draw a straight line to a point that is relatively right x and down y (or left and up if negative values)
func (p *Path) Line(x, y float64) *Path {
	p.b.WriteString(fmt.Sprintf("l %.4f, %.4f ", x, y))
	return p
}

// H x		Draw a line horizontally to the exact coordinate x
func (p *Path) MoveXABS(x float64) *Path {
	p.b.WriteString(fmt.Sprintf("H %.4f ", x))
	return p
}

// h x		Draw a line horizontally relatively to the right x (or to the left if a negative value)
func (p *Path) MoveX(x float64) *Path {
	p.b.WriteString(fmt.Sprintf("h %.4f ", x))
	return p
}

// V y		Draw a line vertically to the exact coordinate y
func (p *Path) MoveYABS(y float64) *Path {
	p.b.WriteString(fmt.Sprintf("V %.4f ", y))
	return p
}

// v y		Draw a line vertically relatively down y (or up if a negative value)
func (p *Path) MoveY(y float64) *Path {
	p.b.WriteString(fmt.Sprintf("v %.4f ", y))
	return p
}

// Z		Draw a straight line back to the start of the path
func (p *Path) Connect() *Path {
	p.b.WriteString("Z")
	return p
}

/*
Curves
*/

// C cX1,cY1 cX2,cY2 eX,eY
// Draw a bezier curve based on two bezier control points and end at specified coordinates
func (p *Path) CurveABS(cx1, cy1, cx2, cy2, x2, y2 float64) *Path {
	p.b.WriteString(fmt.Sprintf("C %.4f,%.4f, %.4f,%.4f, %.4f,%.4f", cx1, cy1, cx2, cy2, x2, y2))
	return p
}

// c Same with all relative values
// TODO

// S cX2,cY2 eX,eY
// Basically a C command that assumes the first bezier control point is a reflection of the last bezier point used in the previous S or C command
func (p *Path) SymmetricABS(cx1, cy1, x2, y2 float64) *Path {
	p.b.WriteString(fmt.Sprintf("S %.4f,%.4f, %.4f,%.4f", cx1, cy1, x2, y2))
	return p
}

// s
//	Same with all relative values

// Q cX,cY eX,eY
// Draw a bezier curve based a single bezier control point and end at specified coordinates
func (p *Path) QuadraticABS(cx1, cy1, x2, y2 float64) *Path {
	p.b.WriteString(fmt.Sprintf("Q %.4f,%.4f, %.4f,%.4f", cx1, cy1, x2, y2))
	return p
}

// q
// Same with all relative values
// TODO:

// T eX,eY
// 	Basically a Q command that assumes the first bezier control point is a reflection of the last bezier point used in the previous Q or T command
func (p *Path) QuadraticSmoothABS(x2, y2 float64) *Path {
	p.b.WriteString(fmt.Sprintf("T %.4f,%.4f", x2, y2))
	return p
}

// t
//	Same with all relative values
/*
A rX,rY rotation, arc, sweep, eX,eY
	Draw an arc that is based on the curve an oval makes. First define the width and height of the oval. Then the rotation of the oval. Along with the end point, this makes two possible ovals. So the arc and sweep are either 0 or 1 and determine which oval and which path it will take.
a
	Same with relative values for eX,eY
*/
