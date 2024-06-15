package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("adds 2 + 2 = 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertCorrectMessage(t, sum, expected)
	})

	t.Run("adds 1 + 1 = 2", func(t *testing.T) {
		sum := Add(1, 1)
		expected := 2

		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}
