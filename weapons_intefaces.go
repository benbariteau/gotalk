package main

import "fmt"

//START OMIT
type weapon interface {
	fire()
}

type phaser struct {
	frequency float64
}

func (p phaser) fire() {
	fmt.Printf("Firing phaser at %v GHz\n", p.frequency)
}

type torpedo struct {
	payloadSize int
}

func (t torpedo) fire() {
	fmt.Printf("Firing torpedo with payload size %v\n", t.payloadSize)
}

//END OMIT

//START2 OMIT
func main() {
	weapons := []weapon{
		phaser{47.0},
		torpedo{42},
	}

	for _, w := range weapons {
		w.fire()
		fmt.Println(w.frequency)
	}
}

//END2 OMIT
