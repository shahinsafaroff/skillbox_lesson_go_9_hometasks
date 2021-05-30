package main

import "fmt"

func main() {
type orangeCount int
type appleCount int
type fruitsCount int

oranges := orangeCount(5)
apples := appleCount(7)

fmt.Println(oranges)
fmt.Println(apples)

fmt.Println(fruitsCount(apples) + fruitsCount(oranges))
}
