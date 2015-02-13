package main

import "fmt"

//START OMIT
func main() {
	fmt.Println(transport(doesNothing(reversePolarity, 3, 7)))
}

func transport(i int) int {
	return i + 37
}

func reversePolarity(i, j int) (int, int) {
	return j, i
}

func doesNothing(f func(int, int) (int, int), i, j int) (num int) {
	m, n := f(i, j)
	num = m + n
	return
}

//END OMIT
