package main

import (
	"fmt"
)

//START OMIT
func main() {
	messagesChan := generateMessages()
	for message := range messagesChan {
		fmt.Println(message)
	}
}

func generateMessages() chan string {
	messagesChan := make(chan string)
	go func() {
		messagesChan <- "Space, the final frontier"
		messagesChan <- "These are the voyages of the Starship Enterprise"
		messagesChan <- "Its continuing mission to explore strange new worlds"
		messagesChan <- "To seek out new life and new civilizations"
		messagesChan <- "To bodly go where no one has gone before"
		close(messagesChan)
	}()
	return messagesChan
}

//END OMIT
