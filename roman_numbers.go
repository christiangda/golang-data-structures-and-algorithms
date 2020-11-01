package interview

// RomanToInt ...
// {given: "I", want: 1},
// {given: "V", want: 5},
// {given: "X", want: 10},
// {given: "L", want: 50},
// {given: "C", want: 100},
// {given: "D", want: 500},
// {given: "M", want: 1000},
func RomanToInt(roman string) int {

	var out int
	R := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	runes := []rune(roman)
	for i := 0; i < len(runes); i++ {

		if i > 0 && R[string(runes[i])] > R[string(runes[i-1])] {
			out += R[string(runes[i])] - 2*R[string(runes[i-1])]
		} else {
			out += R[string(runes[i])]
		}
	}

	return out
}

// IntToRoman ...
func IntToRoman(num int) string {

	// units := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	// tens := []string{"", "X", "XX", "XXX", "XL", "L", "LXX", "LXXX", "XC"}
	// hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	// thousands := []string{"", "M", "MM", "MMM"}

	// out := thousands[num/1000] + hundreds[(num%1000)/100] + tens[(num%100)/10] + units[num%10]

	var out string
	n := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	i := len(n) - 1
	for num > 0 {
		div := num / n[i]
		num %= n[i]
		for div > 0 {
			out = s[i]
			div--
		}
		i--
	}

	return out
}
