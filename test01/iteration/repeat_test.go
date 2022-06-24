package iteration

import (
	"fmt"
	"reflect"
	"testing"
)

//测试
func TestRepeat(t *testing.T) {

	repeated := Repeat("a")
	expected := "aaaaa"
	if !reflect.DeepEqual(repeated, expected) {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	sum := Repeat("a")
	fmt.Println(sum)
	// Output: aaaaa
}
