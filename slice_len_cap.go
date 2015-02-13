package main

import (
	"fmt"
)

func main() {
	// START OMIT
	slice := make([]int, 0, 10)
	fmt.Println("Slice has length", len(slice), "and capacity", cap(slice))
	// END OMIT
}
