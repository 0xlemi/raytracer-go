package canvas

import (
	"fmt"
	"strings"

	"github.com/lemidev/raytracer/pkg/tuple"
)

type Canvas struct {
	Width, Height uint16
	Pixels        [][]tuple.Color
}

func NewCanvas(width, height uint16) *Canvas {
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

func (c *Canvas) PixelAt(x, y uint16) tuple.Color {
	return c.Pixels[x][y]
}

func (c *Canvas) WritePixel(x uint16, y uint16, color tuple.Color) {
	c.Pixels[x][y] = color
}

func (c *Canvas) WriteAllPixels(color tuple.Color) {
	for x, row := range c.Pixels {
		for y := range row {
			c.Pixels[x][y] = color
		}
	}
}

func (c *Canvas) ToPPM() string {
	var s strings.Builder
	// Headers
	s.WriteString("P3\n")
	s.WriteString(fmt.Sprintf("%d %d", c.Width, c.Height))
	s.WriteString("\n")
	s.WriteString("255\n")

	// lineBreakCounter := 0
	var color tuple.Color
	var colorString string
	for i := uint16(0); i < c.Height; i++ {
		for e := uint16(0); e < c.Width; e++ {
			// So it nevers go over 70 Characters each line
			// ((colors(5 * 3) * 3 (if all are 3 digit number)) + whiteSpace(5 * 3)
			// max posible characters = 60
			// if lineBreakCounter >= 5 {
			// 	// s.WriteString("\n")
			// 	lineBreakCounter = 0
			// }
			// lineBreakCounter++

			color = c.PixelAt(e, i)
			colorString = fmt.Sprintf("%d %d %d ",
				uint8(color.Red),
				uint8(color.Green),
				uint8(color.Blue),
			)
			s.WriteString(colorString)
		}
		s.WriteString("\n")
	}
	return s.String()
}

// Show the canvas in console, in readable way
func (c Canvas) Print() {
	for i := 0; i < len(c.Pixels); i++ {
		fmt.Println(c.Pixels[i])
	}
}
