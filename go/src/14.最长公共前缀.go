/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
  if len(strs) == 0 {
    return ""
  }
  
  var res string
  var cur byte
  index := 0
  for {
    if len(strs[0]) <= index {
      return res
    }
    cur = strs[0][index]
    
    for _, v := range strs {
      if len(v) <= index {
        return res
      }
      if v[index] != cur {
        return res
      }
    }
    index++
    res += string(cur)
  }
  return res
}
// @lc code=end

