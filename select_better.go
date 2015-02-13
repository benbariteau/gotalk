package main

import "fmt"

//START OMIT
func main() {
	sevens, nines := getN(7), getN(9)
	sevensOpen, ninesOpen := true, true
	for sevensOpen || ninesOpen {
		var i int
		select {
		case i, sevensOpen = <-sevens:
			fmt.Println(i)
		case i, ninesOpen = <-nines:
			fmt.Println(i)
		}
	}
}

func getN(i int) chan int {
	ch := make(chan int)
	go func() {
		for count := 0; count < 5; count++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

//END OMIT
