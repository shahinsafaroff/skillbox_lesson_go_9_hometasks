package main

import (
	"fmt"
	"math"
)

func main() {
	var firstNumber, secondNumber int16
	fmt.Printf("Enter the first number: ")
	fmt.Scan(&firstNumber)
	fmt.Printf("Enter the second number: ")
	fmt.Scan(&secondNumber)
	result := int64(firstNumber) * int64(secondNumber)
	switch {
	case result < math.MinInt16:
		fmt.Printf("Result of multiplication: %d\nSuitable format: %T\n", result, int32(result))
	case result < math.MinInt8:
		fmt.Printf("Result of multiplication:  %d\nSuitable format: %T\n", result, int16(result))
	case result < 0:
		fmt.Printf("Result of multiplication: %d\nSuitable format: %T\n", result, int8(result))
	case result <= math.MaxUint8:
		fmt.Printf("Result of multiplication: %d\nSuitable format: %T\n", result, uint8(result))
	case result <= math.MaxUint16:
		fmt.Printf("Result of multiplication: %d\nSuitable format: %T\n", result, uint16(result))
	default:
		fmt.Printf("Result of multiplication: %d\nSuitable format: %T\n", result, uint32(result))
	}
}