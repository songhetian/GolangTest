package hello

import "testing"

func TestHello(t *testing.T) {

	got := Hello("chris")

	want := "Hello, chris1"

	if got != want {

		t.Errorf("got %v, want %v", got, want)
	}
}
