package main

import (
	"os"

	"github.com/vatine/3dmandel/pkg/coords"
	"github.com/vatine/3dmandel/pkg/mandel"
)

func main() {
	v := mandel.Render(coords.Coord{-2.0, -2.0, -2.0}, 100, 4.0, 1.0/32.0)
	v.Render(os.Stdout)
}
