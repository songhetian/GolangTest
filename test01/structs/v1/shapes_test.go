package v1

import (
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
}
