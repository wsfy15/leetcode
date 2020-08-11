/*
 * @lc app=leetcode.cn id=336 lang=golang
 *
 * [336] 回文对
 */

// @lc code=start
// 用每个单词的 逆序 构建字典树
// 在每个节点上维护以该节点为终止节点的单词
// 在若在该节点之和的部分能构成回文对，则加入suffixs列表中。
type Node struct {
	children [26]*Node
	words    []int // 以该节点为终止节点的单词的index
	suffix   []int // 保存满足该节点后续字符 为回文字符串 的字符串index
}

func palindromePairs(words []string) [][]int {
	node := &Node{}
	for index, word := range words {
		// 逆序
		bytes := []byte(word)
		i, j := 0, len(bytes)-1
		for i < j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		}
		word = string(bytes)

		cur := node
		if isPalindrome(word) {
			cur.suffix = append(cur.suffix, index)
		}
		for i := 0; i < len(word); i++ {
			ch := word[i]
			if cur.children[ch-'a'] == nil {
				cur.children[ch-'a'] = &Node{}
			}
			cur = cur.children[ch-'a']
			if isPalindrome(word[i+1:]) {
				cur.suffix = append(cur.suffix, index)
			}
		}
		cur.words = append(cur.words, index)
	}

	res := [][]int{}
	for index, word := range words {
		cur := node
		j := 0
		for ; j < len(word); j++ {
			if isPalindrome((word[j:])) {
				for _, anotherIndex := range cur.words {
					if anotherIndex != index {
						res = append(res, []int{index, anotherIndex})
					}
				}
			}

			ch := word[j]
			if cur.children[ch-'a'] == nil {
				break
			}
			cur = cur.children[ch-'a']
		}

		if j == len(word) {
			for _, anotherIndex := range cur.suffix {
				if anotherIndex != index {
					res = append(res, []int{index, anotherIndex})
				}
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end
