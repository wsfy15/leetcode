/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

// @lc code=start
type pair struct {
	num  int
	ones int
}

func sortByBits(arr []int) []int {
	n := len(arr)
	pairs := make([]pair, n)
	for index, num := range arr {
		pairs[index] = pair{num: num, ones: getOnes(num)}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].ones == pairs[j].ones {
			return pairs[i].num < pairs[j].num
		}
		return pairs[i].ones < pairs[j].ones
	})

	for i := 0; i < n; i++ {
		arr[i] = pairs[i].num
	}
	return arr
}

func getOnes(n int) int {
	res := 0
	for n > 0 {
		if n&1 == 1 {
			res++
		}
		n >>= 1
	}
	return res
}

// @lc code=end

