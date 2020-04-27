/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	size, start, max := len(s), 0, 0
	counter := make(map[byte]int) // 记录字符下标
	
	for i := 0; i < size; i++ {
		if v, ok := counter[s[i]]; ok {
			if i - start > max {
				max = i - start
			}
      if v >= start {
        start = v + 1
      }
		}
    counter[s[i]] = i
	}

  if size - start > max {
    max = size - start
  }
	return max
}
// @lc code=end

