/*
 * @lc app=leetcode.cn id=836 lang=golang
 *
 * [836] 矩形重叠
 */

// @lc code=start
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[0] > rec2[0] { // 保证rec1在左边
			rec1, rec2 = rec2, rec1
	}

	if rec2[0] < rec1[2] {
			if rec1[1] > rec2[1] { // 保证rec1在下边
					rec1, rec2 = rec2, rec1
			}
			return rec2[1] < rec1[3] 
	}
	return false
}
// @lc code=end

