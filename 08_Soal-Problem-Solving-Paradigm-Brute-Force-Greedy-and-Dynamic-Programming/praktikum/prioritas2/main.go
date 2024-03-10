package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	cache := make([]int, len(jumps))
	cache[0] = 0

	if len(jumps) > 1 {
		cache[1] = Kurang(jumps[0], jumps[1])
	}

	for i := 2; i < len(jumps); i++ {
		cache[i] = Minimum(cache[i-1]+Kurang(jumps[i], jumps[i-1]), cache[i-2]+Kurang(jumps[i], jumps[i-2]))

		// kira2 gini kak, (batu i + batu i+1) (batu i + batu i+2)  trus dicari minimumnya
	}

	return cache[len(jumps)-1]
}

func Kurang(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func Minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

// Ada katak yang awalnya berada di atas Batu 1. Dia akan mengulangi tindakan berikut beberapa kali untuk mencapai batu N. Jika katak sedang berada di Batu i, lompat ke Batu i + 1 atau Batu i + 2. Di sini, biaya | hi - hj | terjadi, di mana j adalah batu untuk mendarat. Temukan biaya total minimum yang mungkin dikeluarkan sebelum katak mencapai Batu N.
