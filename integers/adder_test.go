package integers

import (
	"fmt"
	"testing"
)

func Exampleadd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {

	assertSame := func(t *testing.T, got int, want int) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
	}

	t.Run("should add two integers", func(t *testing.T) {
		sum := add(2, 2)
		expected := 4

		assertSame(t, sum, expected)
	})
}
