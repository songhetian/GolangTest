package array_silce

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	// numbers := [5]int{1, 2, 3, 4, 5}

	// got := Sum(numbers)
	// want := 15

	t.Run("five", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("three", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}

	Sum(numbers)

	//Output :15
}

func TestSumAll(t *testing.T) {

	t.Run("All/silce", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{0, 9})

		want := []int{3, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d ", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkNumber := func(t *testing.T, got, want []int) {

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d ", got, want)
		}
	}
	t.Run("tail", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9})

		want := []int{2, 9}

		checkNumber(t, got, want)

	})

	t.Run("empty tail", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{3, 4, 5})

		want := []int{0, 9}

		checkNumber(t, got, want)
	})
}
