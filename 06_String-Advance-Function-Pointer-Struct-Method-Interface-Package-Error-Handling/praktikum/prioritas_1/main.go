package main

import (
	"fmt"
	"praktikum/prioritas_1/task"
)

func main() {
	fmt.Println("Task 01")
	fmt.Println(task.PickVocals("alterra academy"))     // aea aae
	fmt.Println(task.PickVocals("sepulsa mantap jiwa")) // eua aa ia
	fmt.Println(task.PickVocals("kopi susu"))           // oi uu

	fmt.Println("\nTask 02")
	fmt.Println(task.SpinString("alta"))    // aalt
	fmt.Println(task.SpinString("alterra")) // aalrtre
	fmt.Println(task.SpinString("sepulsa")) // saesplu

	fmt.Println("\nTask 03")
	fmt.Println(task.GroupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
	fmt.Println(task.GroupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(task.GroupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}
