package main

import (
	"fmt"
	"math"
)

func HitungMinibus(n int, anggota []int) int {
	if len(anggota) != n {
		fmt.Println("Error: Input must be equal with count of family")
		return 0
	}

	count4 := 0
	count3 := 0
	count2 := 0
	count1 := 0

	for _, jumlah := range anggota {
		if jumlah == 4 {
			count4++
		} else if jumlah == 3 {
			count3++
		} else if jumlah == 2 {
			count2++
		} else {
			count1++
		}
	}

	bus := 0

	bus = bus + count4

	bus = bus + count3
	if count1 >= count3 {
		count1 = count1 - count3
	} else {
		count1 = 0
	}

	bus = bus + count2/2
	if count2%2 != 0 {
		bus++
		if count1 > 0 {
			count1--
		}
	}

	if count1 > 0 {
		tambahBus := float64(count1) / 2.0
		bus = bus + int(math.Ceil(tambahBus))
	}

	fmt.Printf("Minimum bus: %d\n", bus)
	return bus
}

func main() {
	fmt.Println("Case 1: ")
	HitungMinibus(5, []int{1, 2, 4, 3, 3})

	fmt.Println("Case 2: ")
	HitungMinibus(8, []int{2, 3, 4, 4, 2, 1, 3, 1})

	fmt.Println("Case 3: ")
	HitungMinibus(5, []int{1, 5})
}
