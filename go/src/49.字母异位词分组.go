/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string]int{}
	var res [][]string
	for _, v := range strs {
		str := SortString(v)
		if i, ok := m[str]; ok {
			res[i] = append(res[i], v)
		} else {
			m[str] = len(res)
			res = append(res, []string{v})
		}
	}

	return res
}
// @lc code=end

