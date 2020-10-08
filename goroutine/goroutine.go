package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})

	go fibogo(ch, sig)

	for i := 0; i < 10; i++  {
		fmt.Println( <- ch)
	}

	sig <- struct{}{}
}

func fibogo(ch chan int, sig chan struct{}) {
	a,b := 0,1
	for {
		select {
		case ch <- a:
			a,b = b,a+b
		case <- sig:
			return
		}
	}
	fmt.Println("graceful")
}