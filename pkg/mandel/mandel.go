package mandel

import (
	"github.com/vatine/3dmandel/pkg/coords"
	"github.com/vatine/3dmandel/pkg/volume"
)

func buildTransform(p, last coords.Coord, t coords.Transform) coords.Transform {
	l := p.Len()

	z := p.Cross(last)

	if last.Len() == 0.0 {
		z = coords.Coord{0.0, 0.0, 1.0}
	}
	if z.Len() == 0.0 {
		z = coords.Coord{0.0, 0.0, 1.0}
	}

	negZ := z.Neg()

	if t[2].Sub(z).Abs() > t[2].Sub(negZ).Abs() {
		z = negZ
	}

	y := z.Cross(p)

	return coords.NewTransform(p, y.ScaleTo(l), z.ScaleTo(l))
}

func step(p, c, last coords.Coord, t coords.Transform) (coords.Coord, coords.Transform) {
	if p.Eq(coords.Origin) {
		return c, coords.UnityTransform()
	}

	t = buildTransform(p, last, t)

	intermediate := t.Transform(p)
	return intermediate.Add(c), t
}

func iter(c coords.Coord, iters int) bool {
	p := coords.Origin
	last := coords.Origin
	t := coords.UnityTransform()
	var next coords.Coord
	for n := 0; n < iters; n++ {
		next, t = step(p, c, last, t)
		last = p
		p = next
		if p.Len() >= 10.0 {
			return false
		}
	}

	return true
}

func Render(min coords.Coord, iters int, side, step float64) volume.Volume {
	v := volume.New(min, side, step)

	for dx := 0.0; dx < side; dx += step {
		for dy := 0.0; dy < side; dy += step {
			for dz := 0.0; dz < side; dz += step {
				c := min.Add(coords.Coord{dx, dy, dz})
				if iter(c, iters) {
					v.Set(c)
				}
			}
		}
	}

	return v
}
