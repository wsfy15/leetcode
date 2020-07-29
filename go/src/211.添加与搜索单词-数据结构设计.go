/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
type trie struct {
	children [26]*trie
	ch byte
	isEnd bool
}

type WordDictionary struct {
	trie *trie
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{ trie: &trie{} }
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	root := this.trie
	for i := 0; i < len(word); i++ {
		if root.children[word[i] - 'a'] == nil {
			root.children[word[i] - 'a'] = &trie{
				ch: word[i],
				isEnd: false,
			}
		}
		root = root.children[word[i] - 'a']
	}
	root.isEnd = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return search(word, this.trie, 0)

}

func search(word string, root *trie, index int) bool {
	for i := index; i < len(word); i++ {
		if word[i] == '.' {
			for j := 0; j < 26; j++ {
				if root.children[j] != nil && search(word, root.children[j], i + 1) {
					return true
				}
			}
			return false
		} else {
			if root.children[word[i] - 'a'] == nil {
				return false
			}
			root = root.children[word[i] - 'a']
		}
	}
	return root.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

