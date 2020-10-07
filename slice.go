package main

import (
	"fmt"
	"strings"
)

var input = "abcdefghijk"

func main() {
	inputArr := strings.Split(input, "")

	result := []string{}
	current := ""
	for _, v := range inputArr {
		current += v
		if len(current) == 2 {
			result = append(result, current)
			current = ""
		}
	}

	if current != "" {
		result = append(result, current+"_")
	}

	fmt.Printf("%v\n", result)
}
