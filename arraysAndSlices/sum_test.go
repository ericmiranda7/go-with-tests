package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 4}, []int{2, 1}, []int{8})
	want := []int{6, 3, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkOutput := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("sum all tails with good slice", func(t *testing.T) {
		got := SumAllTails([]int{0, 2}, []int{2, 9})
		want := []int{2, 9}

		checkOutput(t, got, want)
	})

	t.Run("sum all tails with empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 1, 4, 1})
		want := []int{0, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
