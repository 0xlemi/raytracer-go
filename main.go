package main

import (
	"fmt"
	// "os"

	// c "github.com/lemidev/raytracer/pkg/canvas"
	// p "github.com/lemidev/raytracer/pkg/physical"
	// t "github.com/lemidev/raytracer/pkg/tuple"
)

func main() {

	// start := t.NewPoint(0, 1, 0)
	// velocity := t.NewVector(1, 1.8, 0).Normalize()

	// proj := p.NewProjectile(start, velocity)

	// gravity := t.NewVector(0, -0.1, 0)
	// wind := t.NewVector(-0.01, 0, 0)

	// env := p.NewEnvironment(gravity, wind)

	// // canvas := c.NewCanvas(8, 6)
	// // width 450
	// // heigt 250

	// actual := []t.Point{}
	// for 0 < proj.Position.Y {
	// 	actual = append(actual, proj.Position)
	// 	// x := uint16(proj.Position.X)
	// 	// y := (canvas.Height - 1) - uint16(proj.Position.Y)
	// 	// y := uint16(proj.Position.Y)
	// 	// fmt.Println([2]uint16{x, y})
	// 	proj.Tick(env)
	// 	// canvas.WritePixel(x, y, t.NewColor(255, 255, 255))
	// }

	fmt.Print(fmt.Sprint([4][4]float64{}))
	// fmt.Println([4][4]float64{})

	// ppmData := canvas.ToPPM()

	// file, err := os.Create("images/projectile.ppm")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	file.WriteString(ppmData)
	// 	fmt.Println("File Written.")
	// }

	// file.Close()

	// c := canvas.NewCanvas(10, 2)
	// color := t.NewColor(1, 0.8, 0.6).Cap(1).Multi(255)
	// c.WriteAllPixels(color)

	// c.Print()
	// fmt.Println(c.ToPPM())

	// u.R = 0.1
	// u.B = 5.1
	// u.G = 100.1

	// nums := []float64{x.X, x.Y, x.Z}
	// for _, num := range nums {
	// 	fmt.Println(num)
	// }
}
