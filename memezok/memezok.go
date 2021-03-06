package main

import (
	"fmt"
)

func main() {
	var setBill int
	var getBills [4]int
	fmt.Println("Ларек с лимонадом")
	fmt.Printf("Введите 4 купюры покупателей номиналом 5,10 или 20: ")
	for i := 0; i < 4; i++ {
	repeat:
		fmt.Scan(&setBill)
		if setBill%5 > 0 {
			goto repeat
		}
		getBills[i] = setBill
	}
	fmt.Println(getBills)

	fmt.Println(lemonadeChange(getBills))
}

func lemonadeChange(bills [4]int) bool {
	var notes5, notes10, notes20 int
	for _, bill := range bills {
		switch bill {
		case 5:
			notes5++
		case 10:
			notes10++
			if notes5 > 0 {
				notes5--
			} else {
				return false
			}
		case 20:
			notes20++
			if notes10 > 0 && notes5 > 0 {
				notes10--
				notes5--
			} else if notes5 > 2 {
				notes5 -= 3
			} else {
				return false
			}
		default:
			fmt.Println("Фальшивая купюра")
		}
	}
	fmt.Println("notes5:", notes5, "notes10:", notes10, "notes20:", notes20)
	if notes5 < 0 || notes10 < 0 || notes20 < 0 {
		return false
	}
	return true
}