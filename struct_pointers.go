package main

import "fmt"

//START OMIT
type cargo struct {
	stenboltCount int
}

func (c *cargo) sellStenbolt() {
	c.stenboltCount--
}

func (c cargo) String() string {
	return fmt.Sprintf("We have %v crates of self-sealing stenbolts", c.stenboltCount)
}

func main() {
	cargoHold := cargo{10}
	fmt.Println(cargoHold)
	cargoHold.sellStenbolt()
	fmt.Println(cargoHold)
}

//END OMIT
