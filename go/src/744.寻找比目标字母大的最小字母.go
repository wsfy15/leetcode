/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
// ["c","f","j"] "j"
func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters) - 1
	for low < high {
		mid := low + (high - low) >> 1

		if letters[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	if low == len(letters) - 1 && letters[low] <= target {
		return letters[0]
	}

	return letters[low]
}

func nextGreatestLetter2(letters []byte, target byte) byte {
	n := len(letters)
	min := (26 + letters[0] - target) % 26
	var res byte = letters[0]
	for i := 1; i < n; i++ {
		tmp := (26 + letters[i] - target) % 26 
		if tmp < min {
			min = tmp
			res = letters[i]
		}
	}
	return res
}
// @lc code=end

