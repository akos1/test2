package test2

import "testing"

func Test_aplus(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    int
		b    int
		want int
	}{
		{name: "positive numbers", a: 3, b: 4, want: 7},
		{name: "zero values", a: 0, b: 0, want: 0},
		{name: "negative numbers", a: -3, b: -4, want: -7},
		{name: "mixed signs", a: -3, b: 4, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := aplus(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("aplus() = %v, want %v", got, tt.want)
			}
		})
	}
}
