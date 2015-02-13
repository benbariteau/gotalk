package main

import (
	"fmt"
	"time"
)

//START OMIT
func main() {
	go longLife()

	fmt.Println("Live long and prosper")
	time.Sleep(2 * time.Second)
}

func longLife() {
	time.Sleep(time.Second)
	fmt.Println("Peace and long life")
}

//END OMIT
