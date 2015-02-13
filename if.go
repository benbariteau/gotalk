package main

import "fmt"

//START OMIT
func main() {
	num := 47
	if num < 47 {
		fmt.Println("Run level 3 diagnostic")
	} else if num > 47 {
		fmt.Println("Reverse the polarity")
	} else {
		fmt.Println("Bounce a tachyon beam off the deflector dish!")
	}
}

//END OMIT
