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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := aplus(tt.a, tt.b)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("aplus() = %v, want %v", got, tt.want)
			}
		})
	}
}
