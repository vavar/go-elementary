package captcha_test

import (
	"testing"
	"github.com/vavar/go-elementary/captcha"
)

type captchaInput struct {
	Pattern int
	LeftOper int
	Oper int
	RightOper int
}

func TestCaptcha(t *testing.T) {
	testcases := []struct{
		Message string
		Given captchaInput
		Want string
	}{
		{ Message: "pattern 1 for 1+1", Given: captchaInput{ Pattern: 1, LeftOper: 1, Oper: 1, RightOper: 1}, Want: "1 + one" },
		{ Message: "pattern 1 for 2+1", Given: captchaInput{ Pattern: 1, LeftOper: 2, Oper: 1, RightOper: 1}, Want: "2 + one" },
		{ Message: "pattern 1 for 1-1", Given: captchaInput{ Pattern: 1, LeftOper: 1, Oper: 2, RightOper: 1}, Want: "1 - one" },
		{ Message: "pattern 1 for 2-1", Given: captchaInput{ Pattern: 1, LeftOper: 2, Oper: 2, RightOper: 1}, Want: "2 - one" },
	}
	for _, testcase := range testcases {
		t.Run(testcase.Message, func(t *testing.T) {
			given := testcase.Given
			get := captcha.NewCaptcha(given.Pattern, given.LeftOper, given.Oper, given.RightOper).String()
			if testcase.Want != get {
				t.Errorf("given %d want %q but got %q", given, testcase.Want, get)
			}
		})
	}
}