package output_test

import (
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	t.Logf("t.Logf: log in test")
	t.Errorf("t.Errorf: Fail but continue to execute")
	fmt.Printf("fmt.Printf: do not use me in test...\n")
	t.Fatalf("t.Fatalf: stop executing this function!")
	fmt.Println("After Fatalf: I will not be executed...")
}

func TestOther(t *testing.T) {
	t.Logf("I am TestOther and will be executed, whether the last test failed or not")
}
