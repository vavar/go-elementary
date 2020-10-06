package fizzbuzz

import (
	"fmt"
)

// FizzBuzz ...
func FizzBuzz(n int) string {

	if n%15 == 0 {
		return "FizzBuzz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	return fmt.Sprintf("%d", n)
}
