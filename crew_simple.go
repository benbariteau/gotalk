package main

import "fmt"

func main() {
	crew := getCrew()
	for i := 0; i < len(crew); i++ {
		fmt.Print(crew[i], ", ")
	}
	fmt.Println()
}

func getCrew() []string {
	crew := make([]string, 0)
	crew = append(crew, "Picard", "Riker", "Data", "La Forge", "Worf", "Troi", "Dr. Crusher")
	return crew
}
