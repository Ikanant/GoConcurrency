package main

import (
	"fmt"
	"strings"
)

func main() {
	// Goal: Split this phrase up into different words and print them one by one using channels
	phrase := "These is the sentence I want to split"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i := 0; i < len(words); i++ {
		fmt.Println(<-ch + " ")
	}

}
