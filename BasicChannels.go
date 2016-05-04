package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1)

	ch <- "Hello World"

	fmt.Println(<-ch)
}
