package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2, 8}
	got := Perimeter(rectangle)
	want := 20.0

	verifyGotWant(t, rectangle, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{2.0, 8.0}, want: 20},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		verifyGotWant(t, shape, got, want)
	}

	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T) {
			checkArea(t, tc.shape, tc.want)
		})
	}

}

func verifyGotWant(t *testing.T, shape Shape, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("%#v got %.2f, want %.2f", shape, got, want)
	}
}
