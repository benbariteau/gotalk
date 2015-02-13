package main

import "fmt"

//START OMIT
func main() {
	defer func() {
		fmt.Println("Target destroyed, Captain")
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Firing phasers")
	}
}

//END OMIT
