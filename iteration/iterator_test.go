package iteration

import (
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {

	assertSame := func(t *testing.T, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
	}

	t.Run("should repeat characters N times", func(t *testing.T) {
		repeated := repeat("a", 5)
		expected := "aaaaa"

		assertSame(t, repeated, expected)
	})

}
