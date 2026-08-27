package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	assertCorrectMessage(t, repeated, expected)
}

func TestUpper(t *testing.T) {
	lowered := Upper("hello")
	expected := "HELLO"

	assertCorrectMessage(t, lowered, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	repeated := Repeat("67", 2)
	fmt.Println(repeated)
	// Output: 6767
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
