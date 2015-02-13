package main

import "fmt"

// START OMIT
func main() {
	for _, posting := range []string{"Enterprise", "Voyager", "Deep Space 9"} {
		crew := getCrewByPosting(posting)
		commandCrew := crew[:2]
		fmt.Printf("%v's command crew is %v\n", posting, commandCrew)
	}
}

func getCrewByPosting(posting string) []string {
	var crew []string
	switch posting {
	case "Enterprise":
		crew = []string{"Picard", "Riker", "Data", "La Forge", "Worf", "Troi", "Dr. Crusher"}
	case "Voyager":
		crew = []string{"Janeway", "Chakotay", "Tuvok", "Torres", "Paris", "Kim"}
	case "Deep Space 9":
		crew = []string{"Sisko", "Kira", "Dax", "Odo", "Bashir", "O'Brien"}
	}
	return crew
}

//END OMIT
