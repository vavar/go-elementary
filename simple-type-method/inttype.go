package main 

import (
	"fmt"
	"strconv"
)

type Int int

func (i Int) String() string{
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int){
	*i = Int(n)
}

func (i Int) Int() int {
	return int(i)
}

func main() {
	var i Int = 9
	i.Set(1)
	fmt.Printf("int value: %d , string value: %s\n", i.Int(), i.String())
}