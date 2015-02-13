package main

import (
	"fmt"
)

//START OMIT
func main() {
	numPhasers := 5
	for i := 0; i < numPhasers; i++ {
		fmt.Println("Firing phasers!")
	}

	numTorpedoes := 0
	for numTorpedoes < 3 {
		fmt.Println("Firing photon torpedoes!")
		numTorpedoes++
	}
}

//END OMIT
