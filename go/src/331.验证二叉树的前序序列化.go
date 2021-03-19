/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	n := len(preorder)
	nums := 0
	for i := n - 1; i >= 0; i -= 2 {
		if preorder[i] == '#' {
			nums++
		} else {
			for i >= 0 && '0' <= preorder[i] && preorder[i] <= '9' {
				i--
			}
			i++
			
			if nums >= 2 {
				nums--
			} else {
				return false
			}
		}
	}

	return nums == 1
}
// @lc code=end

