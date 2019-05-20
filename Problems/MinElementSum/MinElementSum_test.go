package MinElementSum

import "testing"

func TestGetMinElementSum(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{"All odd", []int{1, 3, 5, 7}, 9},
		{"All even", []int{2, 4, 6, 8}, 20},
		{"Equal odd & even", []int{1, 2, 3, 4}, 0},
		{"More odd", []int{1, 5, 7, 8, 2}, 0},
		{"More even", []int{1, 4, 6, 8}, 10},
	}

	for _, test := range tests {
		got := GetMinElementSum(test.in)
		if got != test.want {
			t.Errorf("(%v,%v) = %v; want %v", test.name, test.in, got, test.want)
		}
	}
}
