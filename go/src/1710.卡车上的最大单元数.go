/*
 * @lc app=leetcode.cn id=1710 lang=golang
 *
 * [1710] 卡车上的最大单元数
 */

// @lc code=start
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	res := 0
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		count := min(boxTypes[i][0], truckSize)
		res += boxTypes[i][1] * count
		truckSize -= count
	}
	return res
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

