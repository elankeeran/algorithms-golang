// Given a 2D board and a list of words from the dictionary, find all words in the board.
// Each word must be constructed from letters of sequentially adjacent cell,
// where "adjacent" cells are those horizontally or vertically neighboring.
// The same letter cell may not be used more than once in a word.
// Note: All inputs are consist of lowercase letters a-z. The values of words are distinct.

package wordsearch

//Trie trie
type Trie struct {
	childs    [26]*Trie
	endOfChar bool
}

func (t *Trie) insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.childs[idx] == nil {
			node.childs[idx] = &Trie{}
		}
		node = node.childs[idx]
	}

	node.endOfChar = true
}

func (t *Trie) startWith(prefix string) *Trie {
	node := t
	lenP := len(prefix)
	for i := 0; i < lenP; i++ {
		idx := prefix[i] - 'a'
		if node.childs[idx] == nil {
			return nil
		}
		node = node.childs[idx]
	}
	return node
}

func findWords(board [][]byte, words []string) []string {
	trie := &Trie{}
	//build trie tree
	for i := 0; i < len(words); i++ {
		trie.insert(words[i])
	}

	result := map[string]bool{}
	//start to bactracing
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			backTrace(&board, x, y, trie, "", &result)
		}
	}

	list := []string{}
	for k := range result {
		list = append(list, k)
	}
	return list
}

func backTrace(board *([][]byte), x, y int, trie *Trie, prefix string, result *map[string]bool) {
	lenX := len(*board)
	lenY := len((*board)[0])

	if x >= 0 && x < lenX && y >= 0 && y < lenY {
		char := (*board)[x][y]
		if char == '#' {
			return
		}
		(*board)[x][y] = '#' //mark visited
		newPrefix := prefix + string(char)
		node := trie.startWith(newPrefix)

		if node != nil {
			backTrace(board, x-1, y, trie, newPrefix, result)
			backTrace(board, x+1, y, trie, newPrefix, result)
			backTrace(board, x, y-1, trie, newPrefix, result)
			backTrace(board, x, y+1, trie, newPrefix, result)

			if node.endOfChar {
				(*result)[newPrefix] = true
			}
		}
		(*board)[x][y] = char

	}
}

func main() {}
