package bin_exp_backoff

import (
	"testing"
	"time"
)

func TestNumberOfSlotTimes(t *testing.T) {
	cases := []struct {
		Try         int
		ExpectedMin int
		ExpectedMax int
	}{
		{0, 0, 0},
		{1, 0, IntegerPower(2, 1)},
		{10, 0, IntegerPower(2, 10)},
	}

	for _, c := range cases {
		for i := 0; i < 100000; i++ {
			number := NumberOfSlotTimes(c.Try)

			if number < c.ExpectedMin {
				t.Fatalf("number %d < expected min %d!", number, c.ExpectedMin)
			}

			if number > c.ExpectedMax {
				t.Fatalf("number %d > expected max %d!", number, c.ExpectedMax)
			}
		}
	}
}

func TestWaitDuration(t *testing.T) {
	cases := []struct {
		Try         int
		SlotTime    time.Duration
		ExpectedMin time.Duration
		ExpectedMax time.Duration
	}{
		{0, 10 * time.Second, 0, 0},
		{1, 10 * time.Second, 0, 10 * time.Second},
		{2, 10 * time.Second, 0, 30 * time.Second},
		{10, 10 * time.Second, 0, 10230 * time.Second},
	}

	for _, c := range cases {
		for i := 0; i < 100000; i++ {
			number := WaitDuration(c.Try, c.SlotTime)

			if number < c.ExpectedMin {
				t.Fatalf("number %d < expected min %d!", number, c.ExpectedMin)
			}

			if number > c.ExpectedMax {
				t.Fatalf("number %d > expected max %d!", number, c.ExpectedMax)
			}
		}
	}
}
