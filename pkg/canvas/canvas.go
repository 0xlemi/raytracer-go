package canvas

import (
	"fmt"

	"github.com/lemidev/raytracer/pkg/tuple"
)

type Canvas struct {
	Width, Height uint8
	Pixels        [][]tuple.Color
}

func NewCanvas(width, height uint8) *Canvas {
	// Set columns count first, or row lenght
	// [0, 0, ...]
	pixels := make([][]tuple.Color, width)

	// Create row count, or column lenght
	// [0, 0, 0]
	// [0, 0, 0]
	// ...
	for i := range pixels {
		pixels[i] = make([]tuple.Color, height)
	}

	return &Canvas{Width: width, Height: height, Pixels: pixels}
}

func (c *Canvas) PixelAt(x, y uint8) tuple.Color {
	return c.Pixels[x][y]
}

func (c *Canvas) WritePixel(x uint8, y uint8, color tuple.Color) {
	c.Pixels[x][y] = color
}

func (c *Canvas) WriteAllPixels(color tuple.Color) {
  for x, row := range c.Pixels{
		for y:= range row {
			c.Pixels[x][y] = color
		}
	}
}

func (c *Canvas) ToPPM() string {
	s := "P3\n" + fmt.Sprintf("%d %d", c.Width, c.Height) + "\n" + "255\n"

	lineBreakCounter := 0
	for _, row := range c.Pixels {
		for _, color := range row {
			// So it nevers go over 70 Characters each line
			// ((colors(5 * 3) * 3 (if all are 3 digit number)) + whiteSpace(5 * 3)
			// max posible characters = 60
			if lineBreakCounter >= 5 {
				s += "\n"
				lineBreakCounter = 0
			}
			lineBreakCounter++
			s += fmt.Sprintf("%d %d %d ",
				uint8(color.Red),
				uint8(color.Green),
				uint8(color.Blue),
			)
		}
	}
	s += "\n"
	return s
}

// Show the canvas in console, in readable way
func (c Canvas) Print() {
	for i := 0; i < len(c.Pixels); i++ {
		fmt.Println(c.Pixels[i])
	}
}
