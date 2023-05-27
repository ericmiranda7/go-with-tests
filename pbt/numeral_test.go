package pbt

import (
	"testing"
)

var cases = []struct {
	Name   string
	Arabic int
	Roman  string
}{
	{
		"1 means I",
		1,
		"I",
	},
	{
		"2 means II",
		2,
		"II",
	},
	{
		"3 means III",
		3,
		"III",
	},
	{
		"4 means IV",
		4,
		"IV",
	},
	{
		"5 means V",
		5,
		"V",
	},
	{
		"6 means VI",
		6,
		"VI",
	},
	{
		"7 means VII",
		7,
		"VII",
	},
	{
		"8 means VIII",
		8,
		"VIII",
	},
	{
		"9 means IX",
		9,
		"IX",
	},
	{
		"10 means X",
		10,
		"X",
	},
	{
		"11 means XI",
		11,
		"XI",
	},
	{
		"19 means XIX",
		19,
		"XIX",
	},
	{
		"25 means XXV",
		25,
		"XXV",
	},
	{
		"30 means XXX",
		30,
		"XXX",
	},
	{
		"38 means XXXIII",
		38,
		"XXXVIII",
	},
	{
		"39 means XXXIX",
		39,
		"XXXIX",
	},
	{
		"40 means XL",
		40,
		"XL",
	},
	{
		"43 means XLIII",
		43,
		"XLIII",
	},
	{
		"69 means LXIX",
		69,
		"LXIX",
	},
	{
		"82 means LXXXII",
		82,
		"LXXXII",
	},
	{"40 gets converted to XL", 40, "XL"},
	{"47 gets converted to XLVII", 47, "XLVII"},
	{"49 gets converted to XLIX", 49, "XLIX"},
	{"50 gets converted to L", 50, "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %q want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d want %d", got, test.Arabic)
			}
		})
	}
}
