package main

import "fmt"

const boilinfF = 212.0

func main() {
	var f = boilinfF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g or %g \n", f, c)
}
