package wordsearch

import (
	"reflect"
	"testing"
)

func TestWordSearch(t *testing.T) {

	checkSums := func(t *testing.T, got, want []string) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test Case 1", func(t *testing.T) {
		board := [][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		}
		words := []string{"oath", "pea", "eat", "rains", "see"}
		got := findWords(board, words)
		want := []string{"oath", "eat"}
		checkSums(t, got, want)
	})

	t.Run("Test Case 2", func(t *testing.T) {
		board := [][]byte{
			{'o', 'a', 'a', 'n', 'a'},
			{'e', 't', 'a', 'e', 'i'},
			{'i', 'h', 'k', 'r', 'e'},
			{'i', 'f', 'l', 'v', 'e'},
			{'r', 'a', 'i', 'n', 's'},
		}
		words := []string{"oath", "pea", "eat", "rains", "see"}
		got := findWords(board, words)
		want := []string{"oath", "eat", "rains", "see"}
		checkSums(t, got, want)
	})
}
