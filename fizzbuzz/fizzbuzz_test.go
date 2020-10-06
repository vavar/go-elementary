package fizzbuzz_test

import (
	"testing"

	"github.com/vavar/go-elementary/fizzbuzz"
)

func TestGivenOne(t *testing.T) {
	given := 1
	want := "1"
	get := fizzbuzz.FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but got %q", given, want, get)
	}
}

func TestGivenTwo(t *testing.T) {
	given := 2
	want := "2"
	get := fizzbuzz.FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but got %q", given, want, get)
	}
}

func TestGivenThree(t *testing.T) {
	given := 3
	want := "Fizz"
	get := fizzbuzz.FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but got %q", given, want, get)
	}
}

func TestGivenFour(t *testing.T) {
	given := 4
	want := "4"
	get := fizzbuzz.FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but got %q", given, want, get)
	}
}
