package captcha

import (
	"strconv"
	"fmt"
)

var operators = []string{"","+","-","*"}
var number = []string{"0","1","2","3","4","5","6","7","8","9"}
var numberInWord = []string{"zero","one","two","three","four","five","six","seven","eight","nine"}

type Operator int
func NewOperator(n int) Operator{
	return Operator(n)
}

func (o Operator) String() string {
	return operators[int(o)]
}

type NumberOperand int
func NewNumberOperand(n int) NumberOperand {
	return NumberOperand(n)
}

func (n NumberOperand) String() string {
	return strconv.Itoa(int(n))
}

type NumberInWordOperand int
func NewNumberInWordOperand(n int) NumberInWordOperand {
	return NumberInWordOperand(n)
}

func (n NumberInWordOperand) String() string {
	return numberInWord[int(n)]
}

func NewLeftOperand(pattern, n int) fmt.Stringer {
	if pattern == 1 {
		return NewNumberOperand(n)
	}
	return NewNumberInWordOperand(n)
}

func NewRightOperand(pattern, n int) fmt.Stringer {
	if pattern == 1 {
		return NewNumberInWordOperand(n)
	}
	return NewNumberOperand(n)
}

type Captcha struct {
	left fmt.Stringer
	oper fmt.Stringer
	right fmt.Stringer
	result int
}

func Answer(left, oper, right int) int {
	if oper == 1 {
		return left + right
	}
	if oper == 2 {
		return left - right
	}
	return left * right 
}

func NewCaptcha(pattern, left, oper, right int) Captcha {
	return Captcha{ 
		left: NewLeftOperand(pattern, left),
		oper: NewOperator(oper),
		right: NewRightOperand(pattern, right),
	}
}

func (c Captcha) String() string {
	return fmt.Sprintf("%s %s %s", c.left, c.oper, c.right)
}



// ============
// type captchaPattern struct {
// 	LeftOperand []string
// 	RightOperand []string
// 	Operator []string
// }

// var patterns = map[int]captchaPattern{
// 	1: captchaPattern{ LeftOperand: number, RightOperand: numberInWord, Operator: operators },
// 	2: captchaPattern{ LeftOperand: numberInWord, RightOperand: number, Operator: operators },
// }

// // Captcha ...
// func Captcha(pattern, leftOperand, operator, rightOperand int) string {
// 	p := patterns[pattern]
// 	return fmt.Sprintf("%s %s %s", p.LeftOperand[leftOperand],p.Operator[operator], p.RightOperand[rightOperand])
// }
