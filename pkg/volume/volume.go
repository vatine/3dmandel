package volume

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	
	"github.com/vatine/3dmandel/pkg/coords"
)

type Volume interface {
	Set(coords.Coord)
	Render(io.Writer)
	IsFull() bool
	IsEmpty() bool
	mini() coords.Coord
}

type empty struct {
	min  coords.Coord
	side float64
}

func (e empty) Set(c coords.Coord) {
	logrus.Error("Unexpect set on empty.")
}

func (e empty) Render(w io.Writer) {
}

func (e empty) IsEmpty() bool {
	return true
}

func (e empty) IsFull() bool {
	return false
}

func (e empty) mini() coords.Coord {
	return e.min
}

func (e empty) String() string {
	return fmt.Sprintf("empty{%s, %f}", e.min, e.side)
}

type full struct {
	min  coords.Coord
	side float64
}

func (f full) Set(c coords.Coord) {
}

func (f full) Render(w io.Writer) {
	fmt.Fprintf(w, "translate(v = [ %f, %f, %f ]) { cube(size = %f, center=true); }\n", f.min[0], f.min[1], f.min[2], f.side)
}

func (f full) IsEmpty() bool {
	return false
}

func (f full) IsFull() bool {
	return true
}

func (f full) mini() coords.Coord {
	return f.min
}

func (e full) String() string {
	return fmt.Sprintf("full{%s, %f}", e.min, e.side)
}


type partial struct {
	min  coords.Coord
	side float64
	subs [2][2][2]Volume
	step float64
}

func (p partial) mini() coords.Coord {
	return p.min
}

func computeSide(side, step float64) float64 {
	for step < side {
		step = 2.0 * step
	}

	return step
}

func New(min coords.Coord, side, step float64) Volume {
	rv := partial{min: min, step: step}
	rv.side = computeSide(side, step)
	half := rv.side / 2.0
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				delta := coords.Coord{float64(x) * half, float64(y) * half, float64(z) * half}
				rv.subs[x][y][z] = empty{min: min.Add(delta), side: half}
			}
		}
	}

	return &rv
}

func (v partial) findSub(c coords.Coord) (int, int, int) {
	x, y, z := 0, 0, 0
	half := v.side / 2.0
	
	if (c[0] >= v.min[0] + half) {
		x++
	}
	if (c[1] >= v.min[1] + half) {
		y++
	}
	if (c[2] >= v.min[2] + half) {
		z++
	}

	return x, y, z
}

func subPartial(min coords.Coord, side, step float64) *partial {
	rv := partial{min: min, step: step, side: side}
	
	half := rv.side / 2.0
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				delta := coords.Coord{float64(x) * half, float64(y) * half, float64(z) * half}
				rv.subs[x][y][z] = empty{min: min.Add(delta), side: half}
			}
		}
	}

	return &rv
}

func (v *partial) Set(c coords.Coord) {
	x, y, z := v.findSub(c)
	half := v.side / 2.0

	logrus.WithFields(logrus.Fields{
		"c": c, "x": x, "y": y, "z": z, "half": half, "min": v.min,
		"side": v.side,
	}).Debug("partial.Set")
	
	switch {
	case half <= v.step:
		v.subs[x][y][z] = full{min: v.subs[x][y][z].mini(), side: half}
	case v.subs[x][y][z].IsEmpty():
		v.subs[x][y][z] = subPartial(v.subs[x][y][z].mini(), half, v.step)
		v.subs[x][y][z].Set(c)
	case v.subs[x][y][z].IsFull():
		return
	default:
		v.subs[x][y][z].Set(c)
		if v.subs[x][y][z].IsFull() {
			v.subs[x][y][z] = full{v.subs[x][y][z].mini(), half}
		}
		
	}
}

func (v partial) IsEmpty() bool {
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				if !v.subs[x][y][z].IsEmpty() {
					return false
				}
			}
		}
	}

	return true
}

func (v partial) IsFull() bool {
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				if !v.subs[x][y][z].IsFull() {
					return false
				}
			}
		}
	}

	return true
}

func (v partial) Render(w io.Writer) {
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			for z := 0; z < 2; z++ {
				v.subs[x][y][z].Render(w)
			}
		}
	}
}

