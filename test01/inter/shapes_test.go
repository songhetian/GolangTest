package inter

import (
	"testing"
)

func TestAre(t *testing.T) {

	// checkArea := func(t *testing.T, shapes Shape, want float64) {

	// 	t.Helper()

	// 	got := shapes.Area()
	// 	if !reflect.DeepEqual(got, want) {

	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }

	//表格驱动测试
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}
}
