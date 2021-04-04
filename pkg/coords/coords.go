package coords

import (
	"fmt"
	"math"
)

const (
	x = 0
	y = 1
	z = 2
)

// Coordinate/vector type (we're not making too much of a distinction
// between them, for good or bad).
type Coord [3]float64

var Origin Coord = Coord{0.0, 0.0, 0.0}

func (c Coord) String() string {
	return fmt.Sprintf("<%f %f %f>", c[0], c[1], c[2])
}

func (c Coord) Neg() Coord {
	return Coord{-c[0], -c[1], -c[2]}
}

func (c Coord) Abs() float64 {
	return math.Sqrt(c[0]*c[0] + c[1]*c[1] + c[2]*c[2])
}

// Cross-multiply two vectors
func (c1 Coord) Cross(c2 Coord) Coord {
	var rv Coord

	rv[0] = (c1[1] * c2[2]) - (c1[2] * c2[1])
	rv[1] = (c1[2] * c2[0]) - (c1[0] * c2[2])
	rv[2] = (c1[0] * c2[1]) - (c1[1] * c2[0])

	return rv
}

func (c1 Coord) Add(c2 Coord) Coord {
	var rv Coord

	rv[0] = c1[0] + c2[0]
	rv[1] = c1[1] + c2[1]
	rv[2] = c1[2] + c2[2]

	return rv
}

func (c1 Coord) Sub(c2 Coord) Coord {
	var rv Coord

	rv[0] = c1[0] - c2[0]
	rv[1] = c1[1] - c2[1]
	rv[2] = c1[2] - c2[2]

	return rv
}

// Return the length of a vector
func (c Coord) Len() float64 {
	var acc float64

	for _, v := range c {
		acc += v * v
	}

	return math.Sqrt(acc)
}

// Return true is c1 designates a point in the axis-parallel box
// defined by c1 and c2.
func (c Coord) InBox(c1, c2 Coord) bool {
	for ix, v := range c {
		if !((c1[ix] <= v) && (v < c2[ix])) {
			return false
		}
	}
	return true
}

func (c Coord) Scale(s float64) Coord {
	var rv Coord

	rv[0] = s * c[0]
	rv[1] = s * c[1]
	rv[2] = s * c[2]

	return rv
}

func (c Coord) ScaleTo(target float64) Coord {
	l := c.Len()

	if l == 0.0 {
		return c
	}

	scale := target / l

	return c.Scale(scale)
}

func (c1 Coord) Eq(c2 Coord) bool {
	for ix, v := range c1 {
		if v != c2[ix] {
			return false
		}
	}
	return true
}

// The 3D rotational transform
type Transform [3]Coord

var ZeroTransform Transform = NewTransform(Origin, Origin, Origin)

func (t Transform) Transform(p Coord) Coord {
	var rv Coord

	for i := 0; i < 3; i++ {
		rv[i] = p[0]*t[0][i] + p[1]*t[1][i] + p[2]*t[2][i]
	}

	return rv
}

func NewTransform(x, y, z Coord) Transform {
	return Transform{x, y, z}
}

func UnityTransform() Transform {
	return NewTransform(Coord{1.0, 0.0, 0.0}, Coord{0.0, 1.0, 0.0}, Coord{0.0, 0.0, 1.0})
}
