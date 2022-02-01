package potato

import "testing"

func TestHello(t *testing.T) {
	got, _ := Hello("orion")
	want := "Hi, orion. Welcome!"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
