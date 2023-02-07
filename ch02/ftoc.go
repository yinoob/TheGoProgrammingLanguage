package main

import "fmt"

func main() {
	const freez, boil = 32.0, 212.0
	fmt.Printf("%g = %g\n", freez, fToc(freez))
	fmt.Printf("%g = %g\n", boil, fToc(boil))
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
