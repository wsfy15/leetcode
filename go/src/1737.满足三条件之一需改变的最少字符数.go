/*
 * @lc app=leetcode.cn id=1737 lang=golang
 *
 * [1737] 满足三条件之一需改变的最少字符数
 */

// @lc code=start
func minCharacters(a string, b string) int {
	aCh, bCh := make([]int, 26), make([]int, 26)
	aSum, bSum := len(a), len(b)
	for _, ch := range a {
		aCh[ch-'a']++
	}
	for _, ch := range b {
		bCh[ch-'a']++
	}

	return min(makeBigger(aCh, bCh, aSum, bSum), makeEqual(aCh, bCh, aSum, bSum))
}

func makeBigger(aCh, bCh []int, aSum, bSum int) int {
	res := math.MaxInt32
	aPrefix, bPrefix := 0, 0
	for i := 0; i < 25; i++ {
		aPrefix += aCh[i]
		bPrefix += bCh[i]
		res = min(res, aPrefix+bSum-bPrefix, bPrefix+aSum-aPrefix)
	}

	return res
}

func makeEqual(aCh, bCh []int, aSum, bSum int) int {
	res := math.MaxInt32
	for i := 0; i < 26; i++ {
		res = min(res, aSum-aCh[i]+bSum-bCh[i])
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

