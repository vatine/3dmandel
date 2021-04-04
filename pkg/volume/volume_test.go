package volume

import (
	"testing"
	
	// "github.com/sirupsen/logrus"

	"github.com/vatine/3dmandel/pkg/coords"
)

func equalFloat(t *testing.T, seen, want float64) {
	if seen != want {
		t.Errorf("Unexpected uneqal, saw %f, wanted %f", seen, want)
	}
}

func TestPartial1(t *testing.T) {
	v := New(coords.Origin, 1.0, 0.5)

	p, ok := v.(*partial)
	if !ok {
		t.Errorf("Expected p to be a partial under the hood.")
		return
	}
	equalFloat(t, p.step, 0.5)
	equalFloat(t, p.side, 1.0)
	if !v.IsEmpty() {
		t.Errorf("Expected v to be empty, it is not.")
	}
	if v.IsFull() {
		t.Errorf("Did not expect v to be full, it is.")
	}

	v.Set(coords.Coord{0.05, 0.05, 0.05})
	if v.IsEmpty() {
		t.Errorf("Did not expect v to be empty, it is.")
	}
	if v.IsFull() {
		t.Errorf("Did not expect v to be full, it is.")
	}

	if !p.subs[0][0][0].IsFull() {
		t.Errorf("We just set this, it should be full.")
	}
}

func TestPartial2(t *testing.T) {
	v := New(coords.Origin, 1.0, 0.25)

	p, ok := v.(*partial)
	if !ok {
		t.Errorf("Expected p to be a partial under the hood.")
		return
	}
	equalFloat(t, p.step, 0.25)
	equalFloat(t, p.side, 1.0)

	v.Set(coords.Coord{0.1, 0.1, 0.1})
	p2, ok := p.subs[0][0][0].(*partial)
	if !ok {
		t.Errorf("Expected  partial, got %s.", p.subs[0][0][0])
		return 
	}
	if p2.IsEmpty() {
		t.Errorf("p2 is empty.")
	}
	if !p2.subs[0][0][0].IsFull() {
		t.Errorf("We expected p2.subs[0][0][0] to be full, it is %s", p2.subs[0][0][0])
	}

	v.Set(coords.Coord{0.1, 0.1, 0.3})
	v.Set(coords.Coord{0.1, 0.3, 0.1})
	v.Set(coords.Coord{0.1, 0.3, 0.3})
	v.Set(coords.Coord{0.3, 0.1, 0.1})
	v.Set(coords.Coord{0.3, 0.1, 0.3})
	v.Set(coords.Coord{0.3, 0.3, 0.1})
	v.Set(coords.Coord{0.3, 0.3, 0.3})
	_, ok = p.subs[0][0][0].(full)
	if !ok {
		t.Errorf("We expected p3 to be full, it is %v", p.subs[0][0][0])
	}

	if !p.subs[0][0][0].IsFull() {
		t.Errorf("Expected p.subs[0][0][0] to be full, it is not.")
	}
}
