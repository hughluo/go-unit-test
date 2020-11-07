package basic_test

import (
	"testing"

	"github.com/hughluo/go-unit-test/basic"
)

func TestAbs(t *testing.T) {
	input := -1
	got := basic.Abs(input)
	want := 1
	if got != want {
		t.Errorf("Abs(%d) = %d; want %d", input, got, want)
	}
}

func TestAbsWrong(t *testing.T) {
	input := -1
	got := basic.AbsWrong(input)
	want := 1
	if got != want {
		t.Errorf("AbsWrong(%d) = %d; want %d", input, got, want)
	}
}
