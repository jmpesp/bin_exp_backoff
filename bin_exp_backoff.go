package bin_exp_backoff

import (
	"math/rand"
	"time"
)

// from http://grokbase.com/t/gg/golang-nuts/139n3edtq3/go-nuts-integer-exponentiation-in-go#201309217r5yg6nvxwufetnotsk4jjgulm
func IntegerPower(a int, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func NumberOfSlotTimes(try int) int {
	// return a random number between 0 and 2**c − 1
	return rand.Intn(IntegerPower(2, try))
}

func WaitDuration(try int, slotTime time.Duration) time.Duration {
	return slotTime * time.Duration(NumberOfSlotTimes(try))
}
