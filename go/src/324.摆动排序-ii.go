/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */

// @lc code=start
func wiggleSort(nums []int)  {
	// 先通过快速选择得到中位数
	// 将小于中位数的放在数组左边，大于中位数的放右边
	// 左右两部分进行反序穿插
	n := len(nums)
	quickSelect(nums, 0, n, n / 2)
	mid := nums[n/2]

	i, j, k := 0, 0, n - 1
	for i < k {
		if nums[i] < mid {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		} else if nums[i] > mid {
			nums[k], nums[i] = nums[i], nums[k]
			k--
		} else {
			i++
		}
	}

	tmp := make([]int, n)
	copy(tmp, nums)
	mid = n / 2 // 左边元素个数
	if n & 1 == 1 { // 总共奇数个元素，则让左边数组多一个元素
		mid++
	}
	for i := 0; i < mid; i++ {
		nums[2 * i] = tmp[mid - 1 - i]
	}
	for i := mid; i < n; i++ {
		nums[2*(i - mid) + 1] = tmp[n - 1 - (i - mid)]
	}
}

func quickSelect(nums []int, start, end, length int) {
	i, j := start, start
	for i < end {
		// 在i=end-1时就会把快排选择的基准值移动到中间去，从而使左边的元素都比基准值小，右边的则大
		if nums[i] <= nums[end - 1] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			i++
		}
	}

	// nums[j-1]就是基准值，基准值左边有j - 1个数
	// 如果j-1==length
	if j - 1 > length { // 左边的元素超过length个，即超过一半
		quickSelect(nums, start, j - 1, length)
	} else if j <= length { // j是包含了基准值的元素个数，仍小于等于length，说明右侧元素数量比左侧多
		quickSelect(nums, j, end, length)
	}
}
// @lc code=end

