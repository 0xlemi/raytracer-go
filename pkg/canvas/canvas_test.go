package canvas

import (
	// "fmt"
	// "math"
	"testing"

	"github.com/lemidev/raytracer/pkg/tuple"
)

func TestCreateCanvas(t *testing.T) {
	expected := [2]uint8{11, 13}
	actual := [2]uint8{(*NewCanvas(11, 13)).Width, (*NewCanvas(11, 13)).Height}

	if actual != expected {
		t.Errorf("TestCreateCanvas test error. Expected %v, got %v", expected, actual)
	}
}

func TestCanvasDefaultColor000(t *testing.T) {
	expected := tuple.NewColor(0, 0, 0)
	actual := (*NewCanvas(10, 10)).Pixels[3][4]

	if actual != expected {
		t.Errorf("TestCanvasDefaultColor000 test error. Expected %v, got %v", expected, actual)
	}
}

func TestWritePixelToCanvas(t *testing.T) {
	expected := tuple.NewColor(1, 0, 0) // red
	canvas := NewCanvas(10, 10)
	canvas.WritePixel(3, 6, tuple.NewColor(1, 0, 0))
	actual := (*canvas).Pixels[3][6]

	if actual != expected {
		t.Errorf("TestWritePixelToCanvas test error. Expected %v, got %v", expected, actual)
	}
}

func TestPixelAt(t *testing.T) {
	expected := tuple.NewColor(0, 1, 0) // green
	canvas := NewCanvas(10, 10)
	canvas.WritePixel(2, 7, tuple.NewColor(0, 1, 0))
	actual := canvas.PixelAt(2, 7)

	if actual != expected {
		t.Errorf("TestPixelAt test error. Expected %v, got %v", expected, actual)
	}
}

func TestToPPM(t *testing.T) {
	expected := "P3\n" +
		"5 3\n" +
		"255\n" +
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n" +
		"0 0 0 0 0 0 0 127 0 0 0 0 0 0 0 \n" +
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255 \n"
	c := NewCanvas(5, 3)
	c.WritePixel(0, 0, tuple.NewColor(1.5, 0, 0).Cap(1).Multi(255))
	c.WritePixel(2, 1, tuple.NewColor(0, 0.5, 0).Cap(1).Multi(255))
	c.WritePixel(4, 2, tuple.NewColor(-0.5, 0, 1).Cap(1).Multi(255))
	actual := c.ToPPM()

	if actual != expected {
		t.Errorf("TestToPPM test error. Expected %v, got %v", expected, actual)
	}
}

func TestToPPMEveryPixel(t *testing.T) {
	expected := "P3\n" +
		"10 2\n" +
		"255\n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 \n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 \n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 \n" +
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 \n"

	c := NewCanvas(10, 2)
	color := tuple.NewColor(1, 0.8, 0.6).Cap(1).Multi(255)
	c.WriteAllPixels(color)
	actual := c.ToPPM()

	if actual != expected {
		t.Errorf("TestToPPMEveryPixel test error. Expected %v, got %v", expected, actual)
	}
}
