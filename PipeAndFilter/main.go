package main

import "fmt"

/*
Informations go throuigh pipes until it fills a filter. That filter will later determine if the
information can pass or not.
*/

func main() {
	ch := make(chan int)
	go generate(ch)

	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)

		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func generate(ch chan int) {
	// Infinite loop
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
