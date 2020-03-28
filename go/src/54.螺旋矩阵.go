/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
  rows := len(matrix)
  if rows == 0 {
    return nil
  }
  cols := len(matrix[0])
  size := rows * cols
  
  dx := []int{0, 1, 0, -1}
  dy := []int{1, 0, -1, 0}
  x, y, k, startRow, startCol, count  := 0, 0, 0, 0, 0, 0
  res := make([]int, size)
  for count < size {
    res[count] = matrix[x][y]
    count++
    
    x += dx[k]
    y += dy[k]
    if y == cols || x == rows || y == startCol - 1 || x == startRow - 1 {
      k = (k + 1) % 4
      if y == cols {
        y--
        startRow++
        x++
      } else if x == rows {
        x--
        y--
        cols--
      } else if y == startCol - 1 {
        y++
        x--
        rows--
      } else if x == startRow - 1 {
        x++
        y++
        startCol++
      }
    } 
  }
  return res
}
// @lc code=end

