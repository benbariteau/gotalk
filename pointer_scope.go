package main

import "fmt"

//START OMIT
func main() {
	answer := getAnswer()
	fmt.Println(*answer)
}

func getAnswer() *int {
	i := 42
	return &i
}

//END OMIT
