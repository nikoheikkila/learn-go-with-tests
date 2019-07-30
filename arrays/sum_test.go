package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertArraySum := func(t *testing.T, got int, want int, given []int) {
		t.Helper()

		if got != want {
			t.Errorf("got %d, expected %d, given %v", got, want, given)
		}
	}

	t.Run("should compute the sum from a slice of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := sum(numbers)
		want := 6

		assertArraySum(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {

	assertSame := func(t *testing.T, got []int, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	}

	t.Run("should return a slice containing totals for each slice", func(t *testing.T) {
		got := sumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertSame(t, got, want)
	})

}

func TestSumAllTails(t *testing.T) {
	assertSame := func(t *testing.T, got []int, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	}

	t.Run("should compute a sum of slice's tails", func(t *testing.T) {
		got := sumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertSame(t, got, want)
	})

	t.Run("should handle empty slice", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertSame(t, got, want)
	})

}
