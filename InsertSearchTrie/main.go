package main

const alphabetSize int = 26

//TrieNode tree
type TrieNode struct {
	children   [alphabetSize]*TrieNode
	endOfWords bool
}

func getNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false

	for i := 0; i < alphabetSize; i++ {
		node.children[i] = nil
	}

	return node
}

func insert(root *TrieNode, key string) {
	temp := root

	for i := 0; i < len(key); i++ {

		index := key[i] - 'a'

		if temp.children[index] == nil {
			temp.children[index] = getNode()
		}
		temp = temp.children[index]
	}

	temp.endOfWords = true
}

func search(root *TrieNode, key string) bool {
	temp := root

	for i := 0; i < len(key); i++ {

		index := key[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			return false
		}
	}

	return (temp != nil && temp.endOfWords)
}

func main() {}
