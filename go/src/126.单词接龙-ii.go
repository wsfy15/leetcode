/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	i, n := 0, len(wordList)
	for ; i < n; i++ {
		if wordList[i] == endWord {
			break
		}
	}
	if i == n {
		return nil
	}

	prev1, prev2 := map[string][]string{}, map[string][]string{} // 记录每个节点的父节点，有可能该节点可以通过多个父节点到达
	set1, set2 := map[string]bool{}, map[string]bool{} // 记录从两端遍历时每一层的节点
	set1[beginWord] = true
	set2[endWord] = true
	res := [][]string{}
	dualBfs(set1, set2, wordList, prev1, prev2, &res)
	for i := 0; i < len(res); i++ {
		if res[i][0] != beginWord {
			reverse(res[i])
		}
	}
	return res
}

// 保证queue1长度更小，每次从节点数更少的一端扩张
func dualBfs(set1, set2 map[string]bool, wordList []string, prev1, prev2 map[string][]string, res *[][]string) {
	if len(set1) == 0 {
		return
	}
	if len(set1) > len(set2) {
		dualBfs(set2, set1, wordList, prev2, prev1, res)
		return
	}

	flag := false
	newSet := map[string]bool{}
	for key, _ := range set1 {
		if _, ok := set2[key]; ok {
			flag = true
			tmp := [][]string{}
			getPath(key, []string{key}, &tmp, prev1)
			
			tmp2 := [][]string{}
			getPath(key, []string{}, &tmp2, prev2)

			for i := 0; i < len(tmp); i++ {
				reverse(tmp[i])
			}
			for i := 0; i < len(tmp); i++ {
				for j := 0; j < len(tmp2); j++ {
					*res = append(*res, append(tmp[i], tmp2[j]...))
				}
			}
		} else {
			for _, word := range wordList {
				if word == key || set1[word] {
					continue
				}
				if parents, ok := prev1[key]; ok && exist(parents, word) {
					continue
				}
				if canChange(key, word) {
					newSet[word] = true
					prev1[word] = append(prev1[word], key)
				}
			}
		}
	}

	if !flag {
		dualBfs(set2, newSet, wordList, prev2, prev1, res)
	}
}

func exist(parents []string, word string) bool {
	for _, parent := range parents {
		if parent == word {
			return true
		}
	}
	return false
}

func canChange(word1, word2 string) bool {
	n, count := len(word1), 0
	for i := 0; i < n; i++ {
		if word1[i] != word2[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}

func getPath(key string, cur []string, res *[][]string, prev map[string][]string ) {
	if parents, ok := prev[key]; ok {
		for _, parent := range parents {
			cur = append(cur, parent)
			getPath(parent, cur, res, prev)
			cur = cur[:len(cur) - 1]
		}
	} else {
		dst := make([]string, len(cur))
		copy(dst, cur)
		*res = append(*res, dst)
	}
}

func reverse(s []string) {
	i, j := 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
// @lc code=end

