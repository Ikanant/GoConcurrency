package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "You know nothing John Snow"

	phraseSlice := strings.Split(phrase, "")

	ch := make(chan string, len(phraseSlice))

	for _, char := range phraseSlice {
		ch <- char
	}

	close(ch)

	for {
		if msg, ok := <-ch; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}
	}

}
