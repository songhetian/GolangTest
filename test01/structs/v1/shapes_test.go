package v1

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

	t.Run("strurt1", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		assert(t, got, want)
	})

	t.Run("area", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0
		assert(t, got, want)
	})
}

func ExampleArea() {

	area := Area(10.0, 10.0)

	fmt.Println(area)

	//Output 100
}
