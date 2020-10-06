package main

import (
	"fmt"
	"strings"

	"github.com/vavar/go-elementary/fizzbuzz"
)

func star(n int) string {
	var s string
	for i := 1; i <= n; i++ {
		s += fmt.Sprintf("%d", i)
	}
	for i := n - 1; i > 0; i-- {
		s += fmt.Sprintf("%d", i)
	}
	return s
}

func diamond(n int) string {
	var s string
	for i := 1; i <= n; i++ {
		body := star(i)
		bar := strings.Repeat(" ", n-i)
		s += fmt.Sprintf("%s%s\n", bar, body)
	}
	for i := n - 1; i > 0; i-- {
		body := star(i)
		bar := strings.Repeat(" ", n-i)
		s += fmt.Sprintf("%s%s\n", bar, body)
	}
	return s
}

func prime(n int) {
	for i := 1; i <= n; i++ {
		var count uint = 0
		for j := i; j > 0; j-- {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Printf("%d ", i)
		}
	}
}

func main() {
	fmt.Println(diamond(3))

	prime(100)
	fmt.Println()

	for i := 1; i <= 30; i++ {
		fmt.Print(fizzbuzz.FizzBuzz(i), " ")
	}
}
