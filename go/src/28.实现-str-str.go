/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
  sLen := len(haystack)
  pLen := len(needle)
  if pLen == 0 {
    return 0
  }
  if sLen == 0 {
    return -1
  }
  
  next := make([]int, pLen)
  getNext(next, needle)
  j := 0 // 模式串指针
  for i := range haystack {
    for j > 0 && needle[j] != haystack[i] {
      j = next[j - 1] + 1
    }
    
    if needle[j] == haystack[i] {
      j++
    }
    
    if j == pLen {
      return i - pLen + 1
    }
  }
  return -1
}

func getNext(next []int, needle string) {
  next[0] = -1
  pLen := len(needle)
  k := -1 // 可匹配最长前缀的最后一个字符的下标
  for i := 1; i < pLen; i++ {
    for k != -1 && needle[k + 1] != needle[i] {
      k = next[k]
    }
    if needle[i] == needle[k + 1] {
      k++
    }
    next[i] = k
  }
}
// @lc code=end

