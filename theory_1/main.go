package main

import (
	"fmt"
	"unsafe"
)

var (
	ToBe bool = false
	MaxInt8 int8 = 1 << 7 -1
	MaxUint8 uint8 = 1 << 8 - 1
	MaxInt16 int16 = 1 << 15 -1
	MaxUint16 uint16 = 1 << 16 - 1
	MaxInt32 int32 = 1 << 31 - 1
	MaxUint32 uint32 = 1 << 32 - 1
	MaxInt64 int64 = 1 << 63 -1
	MaxUint64 uint64 = 1 << 64 - 1
	MaxByte byte = 1 << 8 - 1
	MaxRune rune = 1 << 31 -1
	MaxInt int = 1 << 63 -1

)

func main() {
	fmt.Printf("Type: %T size: %d byte Value: %v\n", ToBe, unsafe.Sizeof(ToBe), ToBe )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxInt8, unsafe.Sizeof(MaxInt8), MaxInt8 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxUint8, unsafe.Sizeof(MaxUint8), MaxUint8 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxInt16, unsafe.Sizeof(MaxInt16), MaxInt16 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxUint16, unsafe.Sizeof(MaxUint16), MaxUint16 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxInt32, unsafe.Sizeof(MaxInt32), MaxInt32 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxUint32, unsafe.Sizeof(MaxUint32), MaxUint32 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxInt64, unsafe.Sizeof(MaxInt64), MaxInt64 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxUint64, unsafe.Sizeof(MaxUint64), MaxUint64 )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxByte, unsafe.Sizeof(MaxByte), MaxByte )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxRune, unsafe.Sizeof(MaxRune), MaxRune )
	fmt.Printf("Type: %T size: %d byte Value: %v\n", MaxInt, unsafe.Sizeof(MaxInt), MaxInt )
}
