package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum with slice", func(t *testing.T) {
		numbers := []int{5, 4, 3}
		got := Sum(numbers)
		want := 12

		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}
