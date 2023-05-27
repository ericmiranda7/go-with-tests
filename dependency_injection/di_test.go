package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	b := &bytes.Buffer{}
	Greet(b, "henlo kylo")

	got := b.String()
	want := "henlo kylo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
