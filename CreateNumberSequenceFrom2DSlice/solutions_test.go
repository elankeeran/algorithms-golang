package mergesortedarrays

import (
	"reflect"
	"testing"
)

//create sequence number from 2d stored array list
func TestSolveMergeSortedArrays(t *testing.T) {

	checkResult := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test Case 1", func(t *testing.T) {

		board := [][]int{
			{2, 3, 3, 7, 8, 9, 12, 18, 99, 111, 119, 300, 400, 500},
			{3, 6, 7, 7, 7, 100, 250},
			{1, 256},
			{0, 1, 2, 2, 3, 4, 4, 5, 8, 8, 9},
		}
		want := []int{0, 1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 5, 6, 7, 7,
			7, 7, 8, 8, 8, 9, 9, 12, 18, 99, 100, 111, 119, 250, 256,
			300, 400, 500}
		got := solveMergeSortedArrays(board)
		checkResult(t, got, want)
	})

	t.Run("Test Case 2", func(t *testing.T) {

		board := [][]int{
			{1, 2, 3},
			{7, 8, 9},
			{4, 5, 6},
		}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		got := solveMergeSortedArrays(board)
		checkResult(t, got, want)
	})
	t.Run("Test Case 3", func(t *testing.T) {

		board := [][]int{
			{1, 2, 3},
		}
		want := []int{1, 2, 3}
		got := solveMergeSortedArrays(board)
		checkResult(t, got, want)
	})

	t.Run("Test Case 3", func(t *testing.T) {

		board := [][]int{
			{31, 32, 41},
			{51},
			{1},
		}
		want := []int{1, 31, 32, 41, 51}
		got := solveMergeSortedArrays(board)
		checkResult(t, got, want)
	})
}
