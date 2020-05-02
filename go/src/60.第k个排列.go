/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	used := make([]bool, n + 1)
	count := make([]int, n + 1)
	count[1] = 1
	for i := 2; i <= n; i++ {
		count[i] = i * count[i - 1]
	}

	sb := strings.Builder{}
	for j := n; j >= 1; j-- {
		cnt := 1
		if j > 1 && k > count[j - 1] {
      cnt = 1 + (k - 1) / count[j - 1]
			k -= (cnt - 1) * count[j - 1]
		}
		for i := 1; i <= n; i++ {
			if !used[i] {
        cnt--
				if cnt == 0 {
					used[i] = true
					sb.WriteByte('0' + byte(i))
				}
				
			}
		}
	}

	return sb.String()
}
// @lc code=end

