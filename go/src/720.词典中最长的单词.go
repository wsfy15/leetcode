/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start

type trie struct {
	ch byte
	isWord bool
	children [26]*trie
}

type TrieTree struct {
	root *trie
}

func (t *TrieTree) Add(word string) {
	root := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if root.children[ch - 'a'] == nil {
			root.children[ch - 'a'] = &trie{}
		}
		root = root.children[ch - 'a']
	}
	root.isWord = true
}

func (t *TrieTree) Valid(word string) bool {
	root := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if root.children[ch - 'a'].isWord == false {
			return false
		}
		root = root.children[ch - 'a']
	}
	return true
}

func longestWord(words []string) string {
	var res string
	maxLen := 0

	t := &TrieTree{ root: &trie{} }
	for _, word := range words {
		t.Add(word)
	}

	for _, word := range words {
		if t.Valid(word) {
			if len(word) > maxLen || (len(word) == maxLen && word < res) {
				res = word
				maxLen = len(word)
			}
		}
	}

	return res
}
// @lc code=end

