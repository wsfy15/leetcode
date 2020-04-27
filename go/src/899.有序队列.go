/*
 * @lc app=leetcode.cn id=899 lang=golang
 *
 * [899] 有序队列
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

func orderlyQueue(S string, K int) string {
	// 当k > 1时，可以交换任意两个字符，得到字典序最小的字符串
	res := S
	if K == 1 {
		size := len(S)
		for i := 0; i < size; i++ {
			S = S[1:] + string(S[0])
			if S < res {
				res = S
			}
		}
	} else {
		res = SortString(S)
	}
	return res
}
// @lc code=end

