package validatesubsequence

import (
	"testing"
)

func TestValidateSubsequence(t *testing.T) {
	checkSequence := func(t *testing.T, got, want bool) {
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test Case 1 Solution 1 Should return true for valid sequence", func(t *testing.T) {
		array := []int{5, 1, 22, 25, 6, -1, 8, 10}
		sequence := []int{1, 6, -1, 10}
		got := Solution1(array, sequence)
		want := true
		checkSequence(t, got, want)
	})

	t.Run("Test Case 2 Solution 2 Should return true for valid sequence", func(t *testing.T) {
		array := []int{5, 1, 22, 25, 6, -1, 8, 10}
		sequence := []int{1, 6, -1, 10}
		got := Solution2(array, sequence)
		want := true
		checkSequence(t, got, want)
	})

	t.Run("Test Case 3 Solution 2 Should return false for invalid sequence", func(t *testing.T) {
		array := []int{5, 1, 22, 25, 6, -1, 8, 10}
		sequence := []int{1, 7, -1, 7}
		got := Solution2(array, sequence)
		want := false
		checkSequence(t, got, want)
	})
}
