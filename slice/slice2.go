package slice

import (
	"fmt"
)

var input2 = "abcdefghijk"

func slice2() {

	temp := input2 + "_"
	result := []string{}
	for len(temp) > 1 {
		result, temp = append(result, temp[:2]), temp[2:]
	}

	fmt.Printf("%v\n", result)
}
