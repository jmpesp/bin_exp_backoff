package main

import (
	"log"
	"time"

	"github.com/jmpesp/bin_exp_backoff"
)

func main() {
	for i := 0; i < 8; i++ {
		log.Printf("try %d", i)
		time.Sleep(bin_exp_backoff.WaitDuration(i, 250*time.Millisecond))
	}
}
