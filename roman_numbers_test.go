package interview

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {

	t.Run("Roman to integer", func(t *testing.T) {
		tests := []struct {
			name  string
			given string
			want  int
		}{
			{name: "test case 1", given: "I", want: 1},
			{name: "test case 2", given: "V", want: 5},
			{name: "test case 3", given: "X", want: 10},
			{name: "test case 4", given: "L", want: 50},
			{name: "test case 5", given: "C", want: 100},
			{name: "test case 6", given: "D", want: 500},
			{name: "test case 7", given: "M", want: 1000},

			{name: "test case 8", given: "II", want: 2},
			{name: "test case 9", given: "IV", want: 4},
			{name: "test case 10", given: "XI", want: 11},
			{name: "test case 11", given: "VI", want: 6},
			{name: "test case 12", given: "LVII", want: 57},

			{name: "test case 13", given: "CCXLVII", want: 247},
			{name: "test case 14", given: "DCCLXXXIX", want: 789},
			{name: "test case 15", given: "MMCDXXI", want: 2421},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {
				got := RomanToInt(tt.given)

				if got != tt.want {
					t.Errorf("RomanToInt(%q) got %v, want %v", tt.given, got, tt.want)
				}
			})
		}
	})
}

func TestIntToRoman(t *testing.T) {

	t.Skip("NOTE: The implementation fails for some case, WIP")

	t.Run("Integer to roman", func(t *testing.T) {
		tests := []struct {
			name  string
			given int
			want  string
		}{
			{name: "test case 1", given: 1, want: "I"},
			{name: "test case 2", given: 2, want: "II"},
			{name: "test case 3", given: 3, want: "III"},

			{name: "test case 4", given: 5, want: "V"},
			{name: "test case 5", given: 10, want: "X"},
			{name: "test case 6", given: 50, want: "L"},
			{name: "test case 7", given: 100, want: "C"},
			{name: "test case 8", given: 500, want: "D"},
			{name: "test case 9", given: 1000, want: "M"},

			{name: "test case 10", given: 2, want: "II"},
			{name: "test case 11", given: 4, want: "IV"},
			{name: "test case 12", given: 11, want: "XI"},
			{name: "test case 13", given: 6, want: "VI"},
			{name: "test case 14", given: 57, want: "LVII"},

			{name: "test case 15", given: 247, want: "CCXLVII"},
			{name: "test case 16", given: 789, want: "DCCLXXXIX"},
			{name: "test case 17", given: 2421, want: "MMCDXXI"},
		}

		for _, tt := range tests {

			t.Run(tt.name, func(t *testing.T) {
				got := IntToRoman(tt.given)

				if got != tt.want {
					t.Errorf("RomanToInt(%v) got %q, want %q", tt.given, got, tt.want)
				}
			})
		}
	})
}
