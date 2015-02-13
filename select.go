package main

import "fmt"

//START OMIT
func main() {
	sevens := getN(7)
	nines := getN(9)
	for {
		select {
		case i := <-sevens:
			fmt.Println(i)
		case i := <-nines:
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
	}()
	return ch
}

//END OMIT
