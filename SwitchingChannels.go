package main

import (
	"fmt"
)

type Message struct {
	to      []string
	from    string
	content string
}

type ErrorMessage struct {
	errorMsg  string
	originMsg Message
}

func main() {
	msgChan := make(chan Message, 1)
	errChan := make(chan ErrorMessage, 1)

	msg := Message{
		to:      []string{"john@snow.com"},
		from:    "jonathan.hdez92@gmail.com",
		content: "Hello there, how are you",
	}

	errMsg := ErrorMessage{
		errorMsg:  "ERROR ERROR ERROR",
		originMsg: Message{},
	}

	msgChan <- msg
	errChan <- errMsg

	select {
	case receivedM := <-msgChan:
		fmt.Println(receivedM)
	case receivedE := <-errChan:
		fmt.Println(receivedE)
	default:
		fmt.Println("NO message received")
	}

}
