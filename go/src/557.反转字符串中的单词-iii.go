/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
  if len(s) == 0 {
    return ""
  }

  sb := []byte(s)
  start := 0
  for sb[start] == ' ' {
    start++
  }
  for i := start; i < len(sb); i++ {
    if sb[i] == ' ' {
      reverse(sb, start, i - 1)
      for sb[i] == ' ' {
        i++
      }
      start = i
    }
  }
  if start < len(sb) - 1{
    reverse(sb, start, len(sb) - 1)
  }
  return string(sb)
}

func reverse(s []byte, start, end int) {
  for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
// @lc code=end

