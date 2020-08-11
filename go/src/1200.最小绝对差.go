/*
 * @lc app=leetcode.cn id=1200 lang=golang
 *
 * [1200] 最小绝对差
 */

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	res := [][]int{}
	diff := math.MaxInt32
	for i := 0; i < len(arr)-1; i++ {
		tmp := arr[i+1] - arr[i]
		if tmp < diff {
			diff = tmp
			res = [][]int{{arr[i], arr[i+1]}}
		} else if tmp == diff {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}

// @lc code=end

