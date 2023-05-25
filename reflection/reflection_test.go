package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Eric"
	var got []string

	s := struct {
		Name string
	}{expected}

	Walk(s, func(str string) {
		got = append(got, str)
	})

	if got[0] != expected {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

}
