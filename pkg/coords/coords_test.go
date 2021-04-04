package coords

import (
	"testing"
)

func TestPoint(t *testing.T) {
	p1 := Coord{1.0, 0.0, 0.0}
	p2 := Coord{0.0, 1.0, 0.0}
	p3 := Coord{0.0, 0.0, 1.0}

	if !p3.Eq(p1.Cross(p2)) {
		seen := p1.Cross(p2)
		t.Errorf("Cross, saw %s, expected %s", seen, p3)
	}
}

func TestTransform(t *testing.T) {
	p1 := Coord{1.0, 0.0, 0.0}
	p2 := Coord{0.0, 1.0, 0.0}
	p3 := Coord{0.0, 0.0, 1.0}

	cases := []struct{
		t Transform
		p Coord
		e Coord
	}{
		{NewTransform(p1, p2, p3), Coord{0.1, 0.2, 0.3}, Coord{0.1, 0.2, 0.3}},
		{NewTransform(p1, p3, p2), Coord{0.1, 0.2, 0.3}, Coord{0.1, 0.3, 0.2}},
		{NewTransform(Coord{1.0, 2.0, 3.0}, Coord{4.0, 5.0, 6.0}, Coord{7.0, 8.0, 9.0}), Coord{1.0, 2.0, 3.0}, Coord{30.0, 36.0, 42.0}},
	}

	for ix, c := range cases {
		seen := c.t.Transform(c.p)
		want := c.e
		if !seen.Eq(want) {
			t.Errorf("Case %d, saw %s, want %s", ix, seen, want)
		}
	}
}
