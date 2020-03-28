/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children [26]*Trie
	ch byte
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	cur := this
	for i := range word {
		if cur.children[int(word[i] - 'a')] == nil {
			cur.children[int(word[i] - 'a')] = &Trie{
				ch: word[i],
			}
		}
		cur = cur.children[int(word[i] - 'a')]
	}
	cur.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := range word {
		if cur.children[int(word[i] - 'a')] == nil {
			return false
		}
		cur = cur.children[int(word[i] - 'a')]
	}
	return cur.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := range prefix {
		if cur.children[int(prefix[i] - 'a')] == nil {
			return false
		}
		cur = cur.children[int(prefix[i] - 'a')]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

