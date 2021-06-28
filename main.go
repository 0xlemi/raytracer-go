package main

import (
	"fmt"

	"github.com/lemidev/raytracer/pkg/canvas"
	t "github.com/lemidev/raytracer/pkg/tuple"
)

func main() {
	// x := set.Vector(3, 2.5, -14)
	// rows := make([]uint8, 100)

	c := canvas.NewCanvas(10, 2)
	color := t.NewColor(1, 0.8, 0.6).Cap(1).Multi(255)
	c.WriteAllPixels(color)

	c.Print()
	fmt.Println(c.ToPPM())

	// u.R = 0.1
	// u.B = 5.1
	// u.G = 100.1

	// nums := []float64{x.X, x.Y, x.Z}
	// for _, num := range nums {
	// 	fmt.Println(num)
	// }
}
