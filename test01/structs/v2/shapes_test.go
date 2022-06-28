package v2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, got, want float64) {

		if !reflect.DeepEqual(got, want) {
			t.Errorf("value: got %.2f, want %.2f", got, want)
		}
	}

	t.Run("struct p", func(t *testing.T) {

		rect := Rectangle{10.0, 10.0}
		got := Perimeter(rect)
		want := 40.0
		assert(t, got, want)
	})

	t.Run("struct area", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Area()
		want := 100.0
		assert(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circles := Circle{10.0}
		got := circles.Area()
		want := 314.1592653589793
		assert(t, got, want)
	})
}

func ExampleArea() {
	rect := Rectangle{10.0, 10.0}
	area := Area(rect)

	fmt.Println(area)

	//Output 100
}
