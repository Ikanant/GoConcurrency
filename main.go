package main

import (
	"time"
)

func main() {

	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 100; i++ {
			println("Hello", i)
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go", i)
			time.Sleep(godur)
		}
	}()

	// Let's put the main "routine" sleep for a bit
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)

}
