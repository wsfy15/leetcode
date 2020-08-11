/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if A > E {
		return computeArea(E, F, G, H, A, B, C, D)
	}

	if C <= E {
		return (C - A) * (D - B) + (G - E) * (H - F)
	}

	width, height := min(C, G) - E, 0
	if D >= H && B <= H {
		height = H - max(B, F)
	} else if H >= D && D >= F {
		height = D - max(B, F)
	}

	return (C - A) * (D - B) + (G - E) * (H - F) - width * height
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

