package table_test

import (
	"testing"

	"github.com/hughluo/go-unit-test/table"
)

func TestMultiply(t *testing.T) {
	cases := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"b is zero", 1, 0, 0},
		{"two negative numbers", -1, 1, -1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := table.Multiply(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Multiple(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
