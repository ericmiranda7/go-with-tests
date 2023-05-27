package pbt

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumeralRules = []RomanNumeral{
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

func ConvertToRoman(number int) string {
	var res strings.Builder

	for _, rn := range romanNumeralRules {
		for number >= rn.Value {
			res.WriteString(rn.Symbol)
			number -= rn.Value
		}
	}

	return res.String()
}
