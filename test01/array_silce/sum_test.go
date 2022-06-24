package array_silce

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func BenchmarkSum(b *testing.B) {
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func ExampleSum() {
	numbers := [5]int{1, 2, 3, 4, 5}

	Sum(numbers)

	//Output :15
}
