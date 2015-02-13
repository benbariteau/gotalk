package main

import "fmt"

//START OMIT
type weapon struct {
	name  string
	power int
}

func (w weapon) fire() {
	fmt.Printf("Firing %v, it did %v damage\n", w.name, w.power)
}

func main() {
	weapons := []weapon{
		weapon{"Phasers", 1},
		weapon{power: 5, name: "Photon torpedoes"},
		weapon{name: "Deflector dish"},
	}

	for _, w := range weapons {
		w.fire()
	}
}

//END OMIT
