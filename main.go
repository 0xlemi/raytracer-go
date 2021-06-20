package main

import (
	"fmt"

	"github.com/lemidev/raytracer/pkg/set"
)

func main() {
	x := set.NewTuple(3, 2.5, -14, 6)

	nums := []float64{x.X, x.Y, x.Z}
	for _, num := range nums {
		fmt.Println(num)
	}
}
