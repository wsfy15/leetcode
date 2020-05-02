/*
 * @lc app=leetcode.cn id=1095 lang=golang
 *
 * [1095] 山脉数组中查找目标值
 */

// @lc code=start
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

//  1. 先找出山顶元素index
//  2. 从左边的升序部分找target，找不到进入第3步
//  3. 在右边的降序部分找target
func findInMountainArray(target int, mountainArr *MountainArray) int {
	high := mountainArr.length() - 1
	topIndex := findTop(mountainArr, high)
	index := findInAsc(target, 0, topIndex, mountainArr)
	if index != -1 {
		return index
	}
	index = findInDesc(target, topIndex + 1, high, mountainArr)
	return index
}

func findInAsc(target, start, end int, mountainArr *MountainArray) int {
	low, high := start, end

	for low < high {
		mid := low + (high - low) >> 1

		if mountainArr.get(mid) >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if mountainArr.get(low) == target {
		return low
	}
	return -1
}

func findInDesc(target,start, end int, mountainArr*MountainArray) int {
	low, high := start, end

	for low < high {
		mid := low + (high - low) >> 1
		
		if mountainArr.get(mid) > target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	if mountainArr.get(low) == target {
		return low
	}
	return -1
}

func findTop(mountainArr *MountainArray, n int) int {
	low, high := 0, n

	for low < high {
		mid := low + (high - low) >> 1

		if mountainArr.get(mid) < mountainArr.get(mid + 1) {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
// @lc code=end

