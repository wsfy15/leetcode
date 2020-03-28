/*
 * @lc app=leetcode.cn id=365 lang=golang
 *
 * [365] 水壶问题
 */

// @lc code=start
// https://www.bilibili.com/video/av46506390?p=4
// 令x、y的最大公约数为m，x、y桶当前乘水a、b升
// 则 m | x (x%m==0), m | y, m | a, m | b 
// 所以x、y、a、b的线性组合也是m的整数倍
// 两个桶的水量和可能是0、x、y、x+y、a+b-x、a+b-y
// 因此，考虑z是否为m的整数倍，且z <= x+y 即可
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if x == 0 {
		return y == z
	}
	if y == 0 {
		return x == z
	}
	if x + y < z {
		return false
	}

	g := gcd(x, y)
	if z % g == 0 {
		return true
	}
	return false
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for x % y != 0 {
		x, y = y, x % y
	}
	return y
}
// @lc code=end

