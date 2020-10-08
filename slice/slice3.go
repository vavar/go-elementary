package slice

import "fmt"

var input3 = "abcdefghijk"

func slice3() {
	ages := []int{45, 87, 33, 20, 18, 46, 70}
	//ages := []int{45}
	top2, ages := ages[:2], ages[2:]
	for _, v := range ages {
		temp := []int{}
		for _, o := range top2 {
			if v > o {
				temp = append(temp, v)
			} else {
				temp = append(temp, o)
			}
		}
		top2 = top2[1:]
	}
	fmt.Println(top2)
}
