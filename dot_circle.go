package main

import (
	"fmt"
	"math"
	"os"

	c "github.com/lemidev/raytracer/pkg/canvas"
	m "github.com/lemidev/raytracer/pkg/matrix"
	t "github.com/lemidev/raytracer/pkg/tuple"
)

func CreateCircle() {

	points := []t.Tuple{}
	for i := 1; i < 13; i++ {
		radian := (float64(i) * math.Pi) / 6.0
		newPoint := t.NewPoint(0, 30, 0)
		points = append(points, m.IDENTITY_MATRIX4.
			RotateZ(radian).
			MultiTL(newPoint).
			ToTuple())
	}

	canvas := c.NewCanvas(100, 100)

	for _, point := range points {
		canvas.WritePixel(uint16(point.X)+50, uint16(point.Y)+50, t.NewColor(255, 255, 255))
	}

	ppmData := canvas.ToPPM()

	file, err := os.Create("images/clock.ppm")

	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(ppmData)
		fmt.Println("File Written.")
	}

	file.Close()

}
