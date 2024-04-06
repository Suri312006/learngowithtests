package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()

	want := 40.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	// Table driven tests

	// this is an anonymous struct
	areaTests := []struct {
		// slice of structs with two fields, shape and want

		shape Shape
		want  float64
		name  string
	}{

		// filling the struct with cases
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
	}

	// iteration over those cases
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v has area %g want %g", tt.shape, got, tt.want)
			}

		})
	}

	// this helper function doesnt care about the shape it is, nor the method
	// its kind of a black box that does its own job, seperate from our implementation
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	//
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	//
	// }
	//
	// t.Run("rectangles", func(t *testing.T) {
	//
	// 	rect := Rectangle{10.0, 10.0}
	// 	checkArea(t, rect, 100.0)
	//
	// })
	// t.Run("circles", func(t *testing.T) {
	//
	// 	circle := Circle{10.0}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

}
