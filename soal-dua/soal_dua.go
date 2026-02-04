package main

import (
	"fmt"
	"math"
)

func HitungMinibus(n int, anggota []int) int {
	// Memvalidasi Input
	if len(anggota) != n {
		fmt.Println("Error: Input must be equal with count of family")
		return 0
	}

	// Hitung jumlah keluarga berdasarkan ukuran anggota
	count4 := 0
	count3 := 0
	count2 := 0
	count1 := 0

	// Masukkan jumlah anggota ke dalam variabel count untuk setiap anggota
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

	// Inisiasi variabel bus untuk kebutuhan bus
	bus := 0

	// Keluarga beranggota 4 orang membutuhkan 1 bus
	bus = bus + count4

	// Keluarga beranggota 3 orang: Butuh 1 bus, tapi bisa menambah 1 keluarga beranggota 1
	bus = bus + count3
	// Jika ada keluarga beranggota 1, akan dimasukkan ke bus dengan keluarga 3 orang
	if count1 >= count3 {
		count1 = count1 - count3
	} else {
		// Jika keluarga beranggota 1 orang habis, di set menjadi 0
		count1 = 0
	}

	// Keluarga 2 orang diprioritaskan dengan sesama keluarga 2 orang
	// Jika ada 2 keluarga beranggota 2 orang, jadi menambah 1 bus (2+2=4 orang, 2 keluarga)
	// 1 = 1 + 2/2
	bus = bus + count2/2

	// Jika sisa keluarga beranggota 2 orang ganjil
	if count2%2 != 0 {
		bus++ // membutuhkan 1 bus baru

		// Menampung satu keluarga beranggota 1 orang
		if count1 > 0 {
			count1--
		}
	}

	// Keluarga isi 1 orang
	// Hanya bisa diisi keluarga 1 orang + keluarga 1 orang
	if count1 > 0 {
		// Pembulatan keatas (math.Ceil)
		// Contoh: sisa 3 keluarga -> (3+1)/2 = 2 bus
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
