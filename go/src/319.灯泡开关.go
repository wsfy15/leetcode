/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */

// @lc code=start
// 翻转奇数次的最后还是亮的，只有平方数的因子数量是奇数个，其他的都是偶数个
// 4的因子数：1、2、4，2X2只算一个2
// 10的因子数：1、10、2、5
func bulbSwitch(n int) int {
	if n < 2 {
		return n
	}
	return int(math.Sqrt(float64(n)))
}

// @lc code=end

