package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	assertCorrectMessage(t, repeated, expected)
}

func TestBuiltInRepeat(t *testing.T) {
	repeated := BuiltInRepeat("a", 6)
	expected := "aaaaaa"

	assertCorrectMessage(t, repeated, expected)
}

func BenchmarkBuiltInRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuiltInRepeat("a", 5)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
}

func ExampleBuiltInRepeat() {
	repeated := BuiltInRepeat("a", 5)
	fmt.Println(repeated)
}

func assertCorrectMessage(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("expected %q got %q", expected, result)
	}
}
