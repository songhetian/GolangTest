package integers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {

	sum := Add(2, 2)
	expected := 4

	if !reflect.DeepEqual(sum, expected) {

		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
