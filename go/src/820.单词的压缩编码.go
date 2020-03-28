/*
 * @lc app=leetcode.cn id=820 lang=golang
 *
 * [820] 单词的压缩编码
 */

// @lc code=start
type trieNode struct {
	ch byte
	children [26]*trieNode
}

func(this *trieNode) Insert(word string) bool {
	cur := this
	flag := false
	for i := len(word) - 1; i >= 0; i-- {
		if cur.children[word[i]-'a'] == nil {
			flag = true
			cur.children[word[i]-'a'] = &trieNode{ch: word[i]}
		}
		cur = cur.children[word[i]-'a']
	}
	return flag
}

func minimumLengthEncoding(words []string) int {
	// 因为以#结束，所以某个单词能与另一个单词共用同一单词的条件是：其中一个单词是另一个的后缀
	// 因此首先根据单词长度排序 然后根据长度从大到小 构造trie树
	// 构造的时候，是从单词最后一个字符往前的顺序
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	count := 0
	root := &trieNode{}
	for _, v := range words {
		if root.Insert(v) {
			count += len(v) + 1
		}
	}

	return count
}
// @lc code=end

