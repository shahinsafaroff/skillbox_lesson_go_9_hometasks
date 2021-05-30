package main
import "fmt"
const (
	twoInZero = 1 << iota
	_
	twoInTwo
	_
	twoInFour
)
func main() {
	fmt.Printf("Two in Zero %d \n", twoInZero)
	fmt.Printf("Two in Two %d \n", twoInTwo)
	fmt.Printf("Two in Two %d \n", twoInFour)
}
