package matrix

import (
	// "fmt"
	// "reflect"
	"math"
	"testing"

	"github.com/lemidev/raytracer/pkg/tuple"
)

func TestMultiTuplePointByTranslation(t *testing.T) {
	expected := tuple.NewPoint(2, 1, 7)
	p := tuple.NewPoint(-3, 4, 5)
	actual := Translation(5, -3, 2).MultiTL(p)

	if actual != expected {
		t.Errorf("TestMultiTuplePointByTranslation test error. Expected %v, got %v", expected, actual)
	}
}

func TestMultiTuplePointByInverseTranslation(t *testing.T) {
	expected := tuple.NewPoint(-8, 7, 3)
	p := tuple.NewPoint(-3, 4, 5)
	invT, _ := Translation(5, -3, 2).Inverse()
	actual := invT.MultiTL(p)

	if actual != expected {
		t.Errorf("TestMultiTuplePointByInverseTranslation test error. Expected %v, got %v", expected, actual)
	}
}

func TestTranslationVector(t *testing.T) {
	expected := tuple.NewVector(-3, 4, 5)
	v := tuple.NewVector(-3, 4, 5)
	actual := Translation(5, -3, 2).MultiTL(v)

	if actual != expected {
		t.Errorf("TestTranslationVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestScalingToPoint(t *testing.T) {
	expected := tuple.NewPoint(-8, 18, 32)
	actual := Scaling(2, 3, 4).MultiTL(tuple.NewPoint(-4, 6, 8))

	if actual != expected {
		t.Errorf("TestScalingToPoint test error. Expected %v, got %v", expected, actual)
	}
}

func TestScalingToVector(t *testing.T) {
	expected := tuple.NewVector(-8, 18, 32)
	actual := Scaling(2, 3, 4).MultiTL(tuple.NewVector(-4, 6, 8))

	if actual != expected {
		t.Errorf("TestScalingToVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestScalingInverseVector(t *testing.T) {
	expected := tuple.NewVector(-2, 2, 2)
	invT, _ := Scaling(2, 3, 4).Inverse()
	actual := invT.MultiTL(tuple.NewVector(-4, 6, 8))

	if actual != expected {
		t.Errorf("TestScalingInverseVector test error. Expected %v, got %v", expected, actual)
	}
}

func TestReflectionPointXAxis(t *testing.T) {
	expected := tuple.NewPoint(-2, 3, 4)
	actual := Scaling(-1, 1, 1).MultiTL(tuple.NewPoint(2, 3, 4))

	if actual != expected {
		t.Errorf("TestReflectionPointXAxis test error. Expected %v, got %v", expected, actual)
	}
}

func TestRotationXAxisHQ(t *testing.T) {
	expected := tuple.NewPoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2)
	actual := RotationX(math.Pi / 4).MultiTL(tuple.NewPoint(0, 1, 0))

	if !actual.ToTuple().Equals(expected.ToTuple()) {
		t.Errorf("TestRotationXAxisHQ test error. Expected %v, got %v", expected, actual)
	}
}

func TestRotationXAxisFQ(t *testing.T) {
	expected := tuple.NewPoint(0, 0, 1)
	actual := RotationX(math.Pi / 2).MultiTL(tuple.NewPoint(0, 1, 0))

	if !actual.ToTuple().Equals(expected.ToTuple()) {
		t.Errorf("TestRotationXAxisFQ test error. Expected %v, got %v", expected, actual)
	}
}

func TestRotationXAxisHQInverse(t *testing.T) {
	expected := tuple.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	p := tuple.NewPoint(0, 1, 0)
	invR, _ := RotationX(math.Pi / 4).Inverse()
	actual := invR.MultiTL(p)

	if !actual.ToTuple().Equals(expected.ToTuple()) {
		t.Errorf("TestRotationXAxisHQInverse test error. Expected %v, got %v", expected, actual)
	}
}

func TestRotationYAxisHQ(t *testing.T) {
	expected := tuple.NewPoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)
	actual := RotationY(math.Pi / 4).MultiTL(tuple.NewPoint(0, 0, 1))

	if !actual.ToTuple().Equals(expected.ToTuple()) {
		t.Errorf("TestRotationYAxisHQ test error. Expected %v, got %v", expected, actual)
	}
}

func TestRotationYAxisFQ(t *testing.T) {
	expected := tuple.NewPoint(1, 0, 0)
	actual := RotationY(math.Pi / 2).MultiTL(tuple.NewPoint(0, 0, 1))

	if !actual.ToTuple().Equals(expected.ToTuple()) {
		t.Errorf("TestRotationYAxisFQ test error. Expected %v, got %v", expected, actual)
	}
}

func TestRotationZAxisHQ(t *testing.T) {
	expected := tuple.NewPoint(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)
	actual := RotationZ(math.Pi / 4).MultiTL(tuple.NewPoint(0, 1, 0))

	if !actual.ToTuple().Equals(expected.ToTuple()) {
		t.Errorf("TestRotationZAxisHQ test error. Expected %v, got %v", expected, actual)
	}
}

func TestRotationZAxisFQ(t *testing.T) {
	expected := tuple.NewPoint(-1, 0, 0)
	actual := RotationZ(math.Pi / 2).MultiTL(tuple.NewPoint(0, 1, 0))

	if !actual.ToTuple().Equals(expected.ToTuple()) {
		t.Errorf("TestRotationZAxisFQ test error. Expected %v, got %v", expected, actual)
	}
}
