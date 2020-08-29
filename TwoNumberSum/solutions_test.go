package main

import (
	"reflect"
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test Case 1 Two Number Solution 1 target sum 10", func(t *testing.T) {
		got := Solution1(10, 3, 5, -4, 8, 11, 1, -1, 6)
		want := []int{11, -1}
		checkSums(t, got, want)
	})

	t.Run("Test Case 2 Two Number Solution 1 target sum -5", func(t *testing.T) {
		got := Solution1(-5, -7, -5, -3, -1, 0, 1, 3, 5, 7)
		want := []int{-5, 0}
		checkSums(t, got, want)
	})

	t.Run("Test Case 3 Two Number Solution 1 sum return empty array", func(t *testing.T) {
		got := Solution1(-5, 0, 1, 3, 5, 7)
		want := []int{}
		checkSums(t, got, want)
	})

	t.Run("Test Case 4 Two Number Solution 2 target Sum 10", func(t *testing.T) {
		got := Solution2([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
		want := []int{11, -1}
		checkSums(t, got, want)
	})

	t.Run("Test Case 5 Two Number Solution 2 target Sum -5", func(t *testing.T) {
		got := Solution2([]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5)
		want := []int{-5, 0}
		checkSums(t, got, want)
	})

	t.Run("Test Case 6 Two Number Solution 2 sum return empty array", func(t *testing.T) {
		got := Solution2([]int{0, 1, 3, 5, 7}, -5)
		want := []int{}
		checkSums(t, got, want)
	})

	t.Run("Test Case 7 Two Number Solution 3 target Sum 10", func(t *testing.T) {
		got := Solution3([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
		want := []int{-1, 11}
		checkSums(t, got, want)
	})

	t.Run("Test Case 8 Two Number Solution 3 target Sum -5", func(t *testing.T) {
		got := Solution3([]int{-7, -5, -3, -1, 0, 1, 3, 5, 7}, -5)
		want := []int{-5, 0}
		checkSums(t, got, want)
	})

	t.Run("Test Case 9 Two Number Solution 3 sum return empty array", func(t *testing.T) {
		got := Solution3([]int{0, 1, 3, 5, 7}, -5)
		want := []int{}
		checkSums(t, got, want)
	})
}
