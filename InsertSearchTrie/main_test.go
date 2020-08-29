package main

import (
	"reflect"
	"testing"
)

func TestInsertSearchTrie(t *testing.T) {

	words := []string{"a", "and", "an", "go", "golang", "man", "mango"}
	root := getNode()

	for i := 0; i < len(words); i++ {
		insert(root, words[i])
	}

	checkSums := func(t *testing.T, got, want bool) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test Case 1", func(t *testing.T) {
		got := search(root, "a")
		want := true
		checkSums(t, got, want)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		got := search(root, "mango")
		want := true
		checkSums(t, got, want)
	})

	t.Run("Test Case 3", func(t *testing.T) {
		got := search(root, "lang")
		want := false
		checkSums(t, got, want)
	})
}
