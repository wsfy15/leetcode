/*
 * @lc app=leetcode.cn id=777 lang=golang
 *
 * [777] 在LR字符串中交换相邻字符
 */

// @lc code=start
// XL -> LX：L只能左移
// RX -> XR：R只能右移
func canTransform(start string, end string) bool {
	// 去掉X后，是否相等
	ns, ne := strings.Builder{}, strings.Builder{}
	n := len(start)
	for i := 0; i < n; i++ {
		if start[i] != 'X' {
			ns.WriteByte(start[i])
		}
		if end[i] != 'X' {
			ne.WriteByte(end[i])
		}
	}
	if ns.String() != ne.String() {
		return false
	}

	// end中的每个L的坐标，必须小于等于start中对应的L的坐标
	// end中的每个R的坐标，必须大于等于start中对应的R的坐标
	startL, startR, endL, endR := []int{}, []int{}, []int{}, []int{}
	for i := 0; i < n; i++ {
		if start[i] == 'L' {
			startL = append(startL, i)
		} else if start[i] == 'R' {
			startR = append(startR, i)
		}

		if end[i] == 'L' {
			endL = append(endL, i)
		} else if end[i] == 'R' {
			endR = append(endR, i)
		}
	}

	for i := 0; i < len(startL); i++ {
		if startL[i] < endL[i] {
			return false
		}
	}

	for i := 0; i < len(startR); i++ {
		if startR[i] > endR[i] {
			return false
		}
	}

	return true
}
// @lc code=end

