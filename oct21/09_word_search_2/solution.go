// https://leetcode.com/problems/word-search-ii/

package wordsearch2

import "fmt"

func Input() ([][]byte, []string) {
	// return [][]byte{
	//   {'o', 'a', 'a', 'n'},
	//   {'e', 't', 'a', 'e'},
	//   {'i', 'h', 'k', 'r'},
	//   {'i', 'f', 'l', 'v'},
	// }, []string{"oath", "pea", "eat", "rain"}
	// return [][]byte{
	//   {'a'},
	// }, []string{"a"}

	return [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain", "hklf", "hf"}
}

type Trie struct {
	child map[byte]*Trie
	word  *string
}

func NewTrie() *Trie {
	return &Trie{
		child: nil,
		word:  nil,
	}
}

func PrintTrie(n *Trie) {
	fmt.Println(n.child)
	fmt.Println()
	for k, v := range n.child {
		fmt.Printf("%c \n", k)
		PrintTrie(v)
	}
}

func (this *Trie) Insert(word string) {
	cur_node := this
	for i := range word {
		c := word[i]
		next_node, ok := cur_node.child[c]
		if !ok {
			next_node = NewTrie()
			if cur_node.child == nil {
				cur_node.child = make(map[byte]*Trie)
			}
			cur_node.child[c] = next_node
		}
		cur_node = next_node
	}
	cur_node.word = &word
}

func (this *Trie) Remove(word string) bool {
	var _delete func(*Trie, string, int) bool
	_delete = func(root *Trie, word string, index int) bool {
		if index == len(word) {
			root.word = nil
		} else {
			c := word[index]
			next_node, ok := root.child[c]
			if ok && _delete(next_node, word, index+1) {
				delete(root.child, c)
			}
		}
		return root.word == nil && len(root.child) == 0
	}
	return _delete(this, word, 0)
}

func check(board [][]byte, root *Trie, row, col int, result *[]string, origin_trie *Trie) {
	if root == nil {
		return
	}

	if root.word != nil {
		*result = append(*result, *root.word)
		origin_trie.Remove(*root.word)
	}

	tmp := board[row][col]
	board[row][col] = 0
	if row > 0 && board[row-1][col] != 0 && root.child[board[row-1][col]] != nil {
		check(board, root.child[board[row-1][col]], row-1, col, result, origin_trie)
	}
	if col > 0 && board[row][col-1] != 0 && root.child[board[row][col-1]] != nil {
		check(board, root.child[board[row][col-1]], row, col-1, result, origin_trie)
	}
	if row+1 < len(board) && board[row+1][col] != 0 && root.child[board[row+1][col]] != nil {
		check(board, root.child[board[row+1][col]], row+1, col, result, origin_trie)
	}
	if col+1 < len(board[0]) && board[row][col+1] != 0 && root.child[board[row][col+1]] != nil {
		check(board, root.child[board[row][col+1]], row, col+1, result, origin_trie)
	}
	board[row][col] = tmp
}

func Run(board [][]byte, words []string) []string {
	root := NewTrie()
	for i := range words {
		root.Insert(words[i])
	}
	res := make([]string, 0)
	nRow, nCol := len(board), len(board[0])
	for i := 0; i < nRow; i++ {
		for j := 0; j < nCol; j++ {
			c := board[i][j]
			next_node := root.child[c]
			if next_node != nil {
				check(board, next_node, i, j, &res, root)
			}
		}
	}
	return res
}
