package main

import (
	"fmt"
)

var input2 = "abcdefghijk"

func main() {

	temp := input2 + "_"
	result := []string{}
	for len(temp) > 1 {
		result, temp = append(result, temp[:2]), temp[2:]
	}

	fmt.Printf("%v\n", result)
}
