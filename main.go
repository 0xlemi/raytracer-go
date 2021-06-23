package main

import (
	"fmt"

	// "github.com/lemidev/raytracer/pkg/set"
)

func main() {
	// x := set.Vector(3, 2.5, -14)
	// rows := make([]uint8, 100)
	pixels := make([][]uint8, 50)
	for i := range pixels {
		pixels[i] = make([]uint8, 50)
	}

	// u := set.Color{}

	// u.R = 0.1
	// u.B = 5.1
	// u.G = 100.1

	
	print2dArray(pixels)

	// nums := []float64{x.X, x.Y, x.Z}
	// for _, num := range nums {
	// 	fmt.Println(num)
	// }
}

func print2dArray(a [][]uint8) { 
	for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
  }
}
