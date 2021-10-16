// https://leetcode.com/problems/implement-trie-prefix-tree/

package trieprefixtree

import "fmt"

type Trie struct {
	child map[byte]*Trie
	end   bool
	char  byte
}

func Constructor() *Trie {
	return &Trie{
		child: nil,
		end:   false,
		char:  0,
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
			next_node = Constructor()
			next_node.char = c
			if cur_node.child == nil {
				cur_node.child = make(map[byte]*Trie)
			}
			cur_node.child[c] = next_node
		}
		cur_node = next_node
	}
	cur_node.end = true
}

func (this *Trie) Remove(word string) {
	candidate_node := this
	candidate_char := word[0]
	prev_node := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		cur_node := prev_node.child[c]
		if cur_node.end && i != len(word)-1 {
			candidate_node = cur_node
			candidate_char = word[i+1]
		}
		prev_node = cur_node
	}
	fmt.Printf("%c %t\n", prev_node.char, prev_node.end)
	fmt.Printf("%c %t\n", candidate_node.char, candidate_node.end)
	fmt.Printf("%c\n", candidate_char)
	if !prev_node.end { //this mean this word not exist in tree
		return
	}
	if prev_node.child == nil { // this mean this node is the last node in chain
		if len(candidate_node.child) > 1 {
			delete(candidate_node.child, candidate_char)
		} else {
			candidate_node.child = nil
		}
	} else {
		prev_node.end = false
	}
}

func (this *Trie) Search(word string) bool {
	cur_node := this
	i := 0
	for ; i < len(word); i++ {
		next_node, ok := cur_node.child[word[i]]
		if !ok {
			break
		}
		cur_node = next_node
	}
	return i == len(word) && cur_node.end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur_node := this
	i := 0
	for ; i < len(prefix); i++ {
		next_node, ok := cur_node.child[prefix[i]]
		if !ok {
			break
		}
		cur_node = next_node
	}
	return i == len(prefix)
}

func Input() int {
	return 0
}

func Run(dump int) int {
	trie := Constructor()
	trie.Insert("hklf")
	trie.Insert("hf")
	// PrintTrie(trie.root)
	fmt.Println(trie.Search("hklf")) // return True
	fmt.Println(trie.Search("hf"))   // return False

	fmt.Println("---------------------------------------------")
	trie.Remove("hklf")
	fmt.Println(trie.Search("hklf")) // return True
	fmt.Println(trie.Search("hf"))   // return False
	return 0
}
