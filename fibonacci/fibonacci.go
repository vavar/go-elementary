package main

import (
	"fmt"
)

func hof(fn func(int), n int){
	fn(n)
}

type fibonacciFunc func (int)

func closure() func() int {
	a,b := 0,1
	return func() int {
		defer func(){ a,b = b,a+b }()
		return a
	}
}

func fibonacci(n int) {
	fn := closure()
	for i:=0; i<n; i++ {
		fmt.Print(fn(),", ")
	}

	fmt.Println("...")
}

func main() {
	// fibonacci(10)
	hof(fibonacci, 15)
	//fibonacciRecurse(10,0,1)
}