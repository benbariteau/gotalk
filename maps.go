package main

import "fmt"

//START OMIT
func main() {
	pips := make(map[string]int)

	pips["Picard"] = 4
	pips["Riker"] = 3
	fmt.Println(pips["Picard"])
	fmt.Println(pips["O'Brien"])
}

//END OMIT
