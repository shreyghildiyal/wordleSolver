package solver_test

import (
	"testing"

	"github.com/shreyghildiyal/wordleSolver/solver"
)

func TestIncrement1(t *testing.T) {
	input := []int{0, 0, 0, 0, 0}
	expectedOutput := []int{1, 0, 0, 0, 0}
	solver.Increment(input, 3)

	for i := 0; i < len(input); i++ {
		if input[i] != expectedOutput[i] {
			t.Error("The Increment was not done correctly")
		}
	}

}

func TestIncrement2(t *testing.T) {
	input := []int{2, 0, 0, 0, 0}
	expectedOutput := []int{0, 1, 0, 0, 0}
	solver.Increment(input, 3)

	for i := 0; i < len(input); i++ {
		if input[i] != expectedOutput[i] {
			t.Error("The Increment was not done correctly")
		}
	}

}

func TestIncrement3(t *testing.T) {
	input := []int{2, 2, 0, 0, 0}
	expectedOutput := []int{0, 0, 1, 0, 0}
	solver.Increment(input, 3)

	for i := 0; i < len(input); i++ {
		if input[i] != expectedOutput[i] {
			t.Error("The Increment was not done correctly")
		}
	}

}

func TestIncrement4(t *testing.T) {
	input := []int{2, 2, 1, 0, 0}
	expectedOutput := []int{0, 0, 2, 0, 0}
	solver.Increment(input, 3)

	for i := 0; i < len(input); i++ {
		if input[i] != expectedOutput[i] {
			t.Error("The Increment was not done correctly")
		}
	}

}

func TestIncrement5(t *testing.T) {
	input := []int{2, 2, 2, 2, 2}
	expectedOutput := []int{0, 0, 0, 0, 0}
	solver.Increment(input, 3)

	for i := 0; i < len(input); i++ {
		if input[i] != expectedOutput[i] {
			t.Error("The Increment was not done correctly")
		}
	}

}
