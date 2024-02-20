package main

import (
	"fmt"
	"strings"
)

func main() {
	// # Soal Prioritas 1

	// Soal 1
	fmt.Println("-----Soal 1-----")
	const phi float64 = 3.14
	var r float64 = 15
	var t float64 = 15

	v := (phi * (r * r)) * t

	fmt.Printf("Volume = %.2f", v)

	// Soal 2
	fmt.Println("\n\n-----Soal 2-----")
	nilai := 84
	switch {
	case (nilai >= 85) && (nilai <= 100):
		fmt.Println("A")
	case (nilai >= 70) && (nilai <= 84):
		fmt.Println("B")
	case (nilai >= 55) && (nilai <= 69):
		fmt.Println("C")
	case (nilai >= 40) && (nilai <= 50):
		fmt.Println("D")
	case (nilai >= 0) && (nilai <= 39):
		fmt.Println("E")
	default:
		fmt.Println("Nilai Invalid")
	}

	// Soal 3
	fmt.Println("\n-----Soal 3-----")
	for i := 1; i <= 100; i++ {
		switch {
		case (i%4 == 0) && (i%7 == 0):
			fmt.Printf("buzz | ")
		case (i%4 == 0):
			fmt.Printf("%d | ", (i * i))
		case (i%7 == 0):
			fmt.Printf("%d | ", (i * i * i))
		default:
			fmt.Printf("%d | ", i)
		}
	}

	// # Soal Prioritas 2
	// Soal 1
	fmt.Println("\n\n-----Soal 1-----")
	x1 := 26

	var pola string
	for i := 1; i <= x1; i++ {
		if x1%i == 0 {
			if i%2 == 0 {
				pola += "I"
			} else if i%2 != 0 {
				pola += "O"
			}
		}
	}
	fmt.Printf("Pola %d = %s", x1, pola)

	// Soal 2
	fmt.Println("\n\n-----Soal 2-----")

	// init
	var skor_akhir int
	var hasil_akhir string

	budget := 51
	waktu_pengerjaan := 10
	tingkat_kesulitan := 3

	// budget
	switch {
	case (budget >= 0) && (budget <= 50):
		skor_akhir += 4
	case (budget >= 51) && (budget <= 80):
		skor_akhir += 3
	case (budget >= 81) && (budget <= 100):
		skor_akhir += 2
	case (budget > 100):
		skor_akhir += 1
	}

	// waktu pengerjaan
	switch {
	case (waktu_pengerjaan >= 0) && (waktu_pengerjaan <= 20):
		skor_akhir += 10
	case (waktu_pengerjaan >= 21) && (waktu_pengerjaan <= 30):
		skor_akhir += 5
	case (waktu_pengerjaan >= 31) && (waktu_pengerjaan <= 50):
		skor_akhir += 2
	case (waktu_pengerjaan > 50):
		skor_akhir += 1
	}

	// tingkat kesulitan
	switch {
	case (tingkat_kesulitan >= 0) && (tingkat_kesulitan <= 3):
		skor_akhir += 10
	case (tingkat_kesulitan >= 4) && (tingkat_kesulitan <= 6):
		skor_akhir += 5
	case (tingkat_kesulitan >= 8) && (tingkat_kesulitan <= 10):
		skor_akhir += 1
	case (tingkat_kesulitan > 10):
		skor_akhir += 0
	}

	// skor akhir
	switch {
	case (skor_akhir >= 17) && (skor_akhir <= 24):
		hasil_akhir = "High"
	case (skor_akhir >= 10) && (skor_akhir <= 16):
		hasil_akhir = "Medium"
	case (skor_akhir >= 3) && (skor_akhir <= 9):
		hasil_akhir = "Low"
	case (skor_akhir <= 2):
		hasil_akhir = "Impossible"
	}

	fmt.Printf("Budget: %d\nWaktu pengerjaan: %d\nTingkat Kesulitan: %d\nHasil Akhir: %s", budget, waktu_pengerjaan, tingkat_kesulitan, hasil_akhir)

	// # Soal Eksplorasi
	// Anagram >> Memiliki Huruf dan Panjang yang sama

	fmt.Println("\n\n-----Soal Eksplorasi-----")
	kata_pertama := "rusak"
	kata_kedua := "kasur"

	var slc []string
	var result string

	if len(kata_pertama) == len(kata_kedua) {
		// for _, v := range kata_pertama {
		// 	slc = append(slc, string(v))
		// }
		slc = strings.Split(kata_pertama, "")

		for _, v := range slc {
			if !strings.Contains(kata_kedua, v) {
				result = "Bukan Anagram"
				break
			} else {
				result = "Anagram"
			}
		}
	}

	fmt.Printf("Kata pertama = %s\nKata kedua = %s\n=> %s", kata_pertama, kata_kedua, result)
}
