package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestReplace(t *testing.T) {
	replaced := Replace("aaaaa", "j")
	expected := "j"

	if replaced != expected {
		t.Errorf("expected %q but got %q", expected, replaced)
	}
}

func TestUpper(t *testing.T) {
	uppered := Upper("missile")
	expected := "MISSILE"

	if uppered != expected {
		t.Errorf("expected %q but got %q", expected, uppered)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("j", 6)
	fmt.Println(repeat)
	// Output: jjjjjj
}

func ExampleReplace() {
	replace := Replace("oink", "meow")
	fmt.Println(replace)
	// Output: meow
}

func ExampleUpper() {
	upper := Upper("meow")
	fmt.Println(upper)
	// Output: MEOW
}
