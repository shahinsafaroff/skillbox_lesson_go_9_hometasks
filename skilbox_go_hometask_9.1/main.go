package main
import (
	"fmt"
	"math"
)
func main() {
	var uint8Counter, uint16Counter int
	for i := 1; i < math.MaxUint32; i++ {
		if uint16(i) == 0 {
			uint8Counter++
		}
		if uint8(i) == 0 {
			uint16Counter++
		}
	}
	fmt.Printf("Quantity of overflow for uint8: %d\nHowever for uint16: %d\n", uint8Counter, uint16Counter)
}