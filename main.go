package main

import (
	"runtime"
	"time"
)

func main() {
	// Allow the app to use up to 2 processors for the application
	runtime.GOMAXPROCS(2)

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
