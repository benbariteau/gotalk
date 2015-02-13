package main

import "fmt"

//START OMIT
func main() {
	generator := getGenerator()
	fmt.Println(generator())
	fmt.Println(generator())
	fmt.Println(generator())
}

func getGenerator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//END OMIT
