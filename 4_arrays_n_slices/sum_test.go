package arraysnslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size ", func(t *testing.T) {

		// this is a slice in the go language, where the length of the array
		// doesnt need to be known
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{1, 2, 3, 4, 5}, []int{10, 2, 1})
	want := []int{3, 15, 13}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v ", got, want)

		}

	}

	t.Run("Make the sums of the tails of non-empty slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2, 3, 4}, []int{2, 3, 4}, []int{10, 1})
		want := []int{10, 9, 11}

		checkSums(t, got, want)

	})

	t.Run("Make the sums of normal slices, and some  empty slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2, 3, 4}, []int{2, 3, 4}, []int{})
		want := []int{10, 9, 0}

		checkSums(t, got, want)
	})
}
