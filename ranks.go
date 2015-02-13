package main

import "fmt"

//START OMIT
func main() {
	ranks := map[string]string{
		"Sisko":   "Commander",
		"Dax":     "Lieutenant",
		"Bashir":  "Lieutenant",
		"O'Brien": "Chief",
	}
	for crewmember, rank := range ranks {
		fmt.Printf("%v's rank is %v\n", crewmember, rank)
	}
	rank, ok := ranks["Kira"]
	if ok {
		fmt.Printf("Kira's rank is %v\n", rank)
	} else {
		fmt.Println("Kira isn't in Starfleet!")
	}
}

//END OMIT
