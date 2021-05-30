package main

import "fmt"

var oranges = 10

var (
	apples = 5
	mandarines = 8
)
const (
	firstSpartan int = 1 + iota
	secondSpartan
	thirdSpartan
	fourthSpartan
)

func main() {
bananas := 9

	fmt.Printf("oranges: %d \n", oranges)
	fmt.Printf("apples: %d \n", apples)
	fmt.Printf("mandarines: %d \n", mandarines)
	fmt.Printf("bananas: %d \n", bananas)
	fmt.Printf("First Spartan: %d \n", firstSpartan)
	fmt.Printf("Second Spartan: %d \n", secondSpartan)
	fmt.Printf("Third Spartan: %d \n", thirdSpartan)
	fmt.Printf("Fourth Spartan: %d \n", fourthSpartan)

}
