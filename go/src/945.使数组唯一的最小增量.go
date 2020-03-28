/*
 * @lc app=leetcode.cn id=945 lang=golang
 *
 * [945] 使数组唯一的最小增量
 */

// @lc code=start
func minIncrementForUnique(A []int) int {
	size := len(A)
	if size < 2 {
		return 0
	}

	sort.Ints(A)
	count := 0
	left := 0
	for i := 1; i < size; i++ {
		if A[i] == A[i - 1] {
			left++
			count -= A[i] // 先减至0，之后找到某个未出现的值时可以直接加
		} else {
			// 找到一部分空缺的数值
			pos := A[i] - A[i - 1] - 1
			if pos == 0 {
				continue
			}

			if left < pos {
				pos = left
			}
			count += pos * A[i-1] + pos*(1+pos)/2
			left -= pos
		}
	}

	if left > 0 {
		count += left * A[size - 1] + left*(left+1)/2
	}

	return count
}

func minIncrementForUnique2(A []int) int {
	size := len(A)
	if size < 2 {
		return 0
	}

	sort.Ints(A)
	count := 0
	for i := 1; i < size; i++ {
		if A[i] == A[i - 1] {
			A[i]++
			count++
		}
		for j := i + 1; j < size && A[j] == A[i-1]; j++ {
			A[j]++
			count++
		}
	}
	return count
}
// @lc code=end

