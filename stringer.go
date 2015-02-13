package main

import "fmt"

type myInt int

func (i myInt) String() string {
	return fmt.Sprint(int(i))
}

func main() {
	var i myInt = 47
	fmt.Print(i.String())
}
