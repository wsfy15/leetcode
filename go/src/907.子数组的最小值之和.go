/*
 * @lc app=leetcode.cn id=907 lang=golang
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
func sumSubarrayMins(A []int) int {
	// 利用单调栈
	mod := 1e9 + 7
	size, res := len(A), 0
	left, right := make([]int, size), make([]int, size)

	// 找到该元素左右两边 不比 该元素大的元素的下标
	stack := []int{}
	for i := 0; i < size; i++ {
		for len(stack) > 0 && A[stack[len(stack) - 1]] >= A[i] {
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack) - 1]
		}

		stack = append(stack, i)
	}

	stack = []int{}
	for i := size - 1; i >= 0; i-- {
		for len(stack) > 0 && A[stack[len(stack) - 1]] > A[i] {
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			right[i] = size
		} else {
			right[i] = stack[len(stack) - 1]
		}

		stack = append(stack, i)
	}

	for i := 0; i < size; i++ {
		res += A[i] * (right[i] - i) * (i - left[i])
	}
	return res % int(mod)
}


func sumSubarrayMins2(A []int) int {
	mod := 1e9 + 7
	size, res := len(A), 0
	left, right := make([]int, size), make([]int, size)

	// 找到该元素左右两边 不比 该元素大的元素的下标
	for i := 0; i < size; i++ {
		j := i - 1
		for j >= 0 && A[j] >= A[i] {
			j--
		} 
		left[i] = j

		j = i + 1
		for j < size && A[j] > A[i] {
			j++
		} 
		right[i] = j
	}	

	for i := 0; i < size; i++ {
		res += A[i] * (right[i] - i) *(i - left[i])
	}
	return res % int(mod)
}
// @lc code=end

