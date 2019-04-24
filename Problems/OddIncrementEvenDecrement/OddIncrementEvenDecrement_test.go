package OddIncrementEvenDecrement

import (
	"fmt"
	"strings"
	"testing"
)

//TestGenerateSlice using Table-Driven Test
func TestGenerateSlice(t *testing.T) {

	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"Zero size input", []int{}, []int{}},
		{"One size input", []int{1}, []int{2}},
		{"Odd size input", []int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 6}},
		{"Odd size input with all zeros", []int{0, 0, 0, 0, 0, 0, 0}, []int{1, -1, 1, -1, 1, -1, 1}},
		{"Even size input", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{"Even size input with all zeros", []int{0, 0, 0, 0, 0, 0}, []int{1, -1, 1, -1, 1, -1}},
	}

	for _, test := range tests {
		got := GenerateSlice(test.in)
		if !areEqual(test.want, got) {
			t.Errorf("%q", test.name)

			for index := range test.in {
				t.Errorf("(%v) = %v; want %v", test.in[index], got[index], test.want[index])
			}
		}
	}
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func areEqual(input []int, output []int) bool {
	if len(input) != len(output) {
		return false
	}

	for i := range input {
		if input[i] != output[i] {
			return false
		}
	}

	return true
}
