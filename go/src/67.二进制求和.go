/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
  if len(a) < len(b) { // 保证a长度不比B短的
    a, b = b, a
  }
  aLen, bLen := len(a), len(b)
  if bLen == 0 {
    return a
  }
  
  res, index := make([]byte, aLen), aLen - 1
  i, j, over := aLen - 1, bLen - 1, 0
  for i >= 0 && j >= 0 {
    tmp := int(a[i] - '0' + b[j] - '0') + over
    res[index] = byte(tmp % 2) + '0'
    over = tmp / 2
    index--
    i--
    j--
  }
  for i >= 0 {
    tmp := int(a[i] - '0') + over
    res[index] = byte(tmp % 2) + '0'
    over = tmp / 2
    index--
    i--
  }
  if over == 1 {
    res = append([]byte{'1'}, res...)
  }
  return string(res)
}
// @lc code=end

