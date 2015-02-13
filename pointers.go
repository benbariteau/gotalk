package main

import "fmt"

//START OMIT
func main() {
	theAnswer := 42

	inflation(&theAnswer)
	fmt.Println(theAnswer)
}

func inflation(i *int) {
	*i += 5
}

//END OMIT
