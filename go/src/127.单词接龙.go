/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 0
	}

	i, n := 0, len(wordList)
	for ; i < n; i++ {
		if wordList[i] == endWord {
			break
		}
	}
	if i == n {
		return 0
	}

	queue1, queue2 := []string{beginWord}, []string{endWord}
	set1, set2 := map[string]bool{ beginWord: true }, map[string]bool{endWord: true} // 记录从两端遍历时每一层的节点
	count := 1
	for {
		if len(queue1) * len(queue2) == 0 {
			return 0
		}
		if len(queue1) > len(queue2) {
			queue1, queue2 = queue2, queue1
			set1, set2 = set2, set1
		}

		size := len(queue1)
		for i := 0; i < size; i++ {
			cur := queue1[i]
			for _, word := range wordList {
				if !set1[word] {
					if canChange(cur, word) {
						if set2[word] {
							return count + 1
						} else {
							set1[word] = true
							queue1 = append(queue1, word)
						}
					}
				}
			}
		}
		queue1 = queue1[size:]
		count++
	}

	return 0
}

// func ladderLength(beginWord string, endWord string, wordList []string) int {
// 	// BFS
// 	queue := []string{beginWord}
// 	visited := map[string]bool{}
// 	visited[beginWord] = true
// 	count := 1
// 	for len(queue) > 0 {
// 		size := len(queue)
// 		for i := 0; i < size; i++ {
// 			word := queue[i]
// 			for _, w := range wordList {
// 				if !visited[w] && canChange(word, w) {
// 					if w == endWord {
// 						return count + 1
// 					}
// 					queue = append(queue, w)
// 					visited[w] = true
// 				}
// 			}
// 		}
// 		count++
// 		queue = queue[size:]
// 	}

// 	return 0
// }

func canChange(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	diff := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return true
}
// @lc code=end

