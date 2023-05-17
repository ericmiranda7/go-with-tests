package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("k", 3)
	want := "kkk"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("bah", 3)
	}
}

func ExampleRepeat() {
	rs := Repeat("ha", 3)
	fmt.Println(rs)
	// Output: hahaha
}
