package main

import "fmt"

func main() {
	var i uint8 = 10
	for ; i != 254; i-- {
		if i == 0 {
			fmt.Println("Coca-cola run out off, but .....")
			fmt.Println("We will start again .....")
			continue
		}
		fmt.Printf("Left bottles %d \n", i)
	}
}
