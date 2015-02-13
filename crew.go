package main

import "fmt"

//START OMIT
func main() {
	for i, member := range getCrew(1) {
		fmt.Printf("The %vth crewmember is %v\n", i, member)
	}
}

func getCrew(season int) []string {
	crew := []string{"Picard", "Riker", "Data", "La Forge", "Worf", "Troi", "Dr. Crusher"}
	return crew
}

//END OMIT
