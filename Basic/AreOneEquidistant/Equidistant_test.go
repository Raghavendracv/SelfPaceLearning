package AreOneEquidistant

import "testing"

//TestAreEquistant using Table-Driven Test
func TestAreEquistant(t *testing.T) {

	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"Equidistant string starts with zero and ends with zero", "0010010010", true},
		{"Equidistant string starts with zero and ends with one", "1001001001", true},
		{"Equidistant string starts with one and ends with zero", "10010010", true},
		{"Equidistant string starts with one and ends with one", "1001001", true},
		{"Equidistant string with all ones", "111111", true},
		{"Non equidistant string stats with zero and ends with zero", "01010101010000010001010101010101010010", false},
		{"Non equidistant string stats with zero and ends with one", "0101010101000001000101010101010101001", false},
		{"Non equidistant string stats with one and ends with zero", "1010101010000010001010101010101010010", false},
		{"Non equidistant string stats with one and ends with one", "101010101000001000101010101010101001", false},
		{"Non equidistant string with all zeros", "00000000", false},
	}

	for _, test := range tests {
		got := AreEquistant(test.in)
		if got != test.want {
			t.Errorf("(%q,%q) = %t; want %t", test.name, test.in, got, test.want)
		}
	}
}
