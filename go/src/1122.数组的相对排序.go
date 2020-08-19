/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	index, n := 0, len(arr1)
	for _, num := range arr2 {
		for i := index; i < n; i++ {
			if arr1[i] == num {
				arr1[index], arr1[i] = arr1[i], arr1[index]
				index++
			}
		}
	}
	sort.Ints(arr1[index:])
	return arr1
}

// @lc code=end

