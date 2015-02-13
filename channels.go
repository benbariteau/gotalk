package main

import "fmt"

//START OMIT
func main() {
	messagesChan := make(chan string)

	go generateMessages(messagesChan)
	for i := 0; i < 5; i++ {
		fmt.Println(<-messagesChan)
	}
}

func generateMessages(messagesChan chan string) {
	messagesChan <- "Space, the final frontier"
	messagesChan <- "These are the voyages of the Starship Enterprise"
	messagesChan <- "Its continuing mission to explore strange new worlds"
	messagesChan <- "To seek out new life and new civilizations"
	messagesChan <- "To bodly go where no one has gone before"
}

//END OMIT
