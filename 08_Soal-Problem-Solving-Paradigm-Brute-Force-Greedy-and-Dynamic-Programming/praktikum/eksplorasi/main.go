package main

import "fmt"

type RomawiSymbol struct {
	Number int
	Symbol string
}

var Symbol = []RomawiSymbol{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {
	fmt.Println(ConvertRomawi(4))
	fmt.Println(ConvertRomawi(9))
	fmt.Println(ConvertRomawi(23))
	fmt.Println(ConvertRomawi(2021))
	fmt.Println(ConvertRomawi(1646))
}

func ConvertRomawi(n int) string {
	var result string

	for n > 0 {
		for _, v := range Symbol {
			for n >= v.Number {
				result += v.Symbol
				n -= v.Number
			}
		}
	}

	return result
}
