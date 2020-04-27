/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
// 二分查找搜索 区间的左边界下标
// 区间左边界下标范围是[0, len(arr) - k]
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr) - k
	for left < right {
		mid := left + (right - left) >> 1

		// 这里的比较要用一个大小为k+1的区间，用大小为k的区间比较没意义
		// if abs(arr[mid] - x) <= abs(arr[mid + k] - x) {
		// 考虑两个极端情况，x在区间左边和区间右边
		// 以及x的值在区间范围内 
		// 因此可以不使用abs函数 
		if x - arr[mid] <= arr[mid + k] -x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return arr[left: left + k]
}

// 双指针从前后删除元素
func findClosestElements3(arr []int, k int, x int) []int {
	n := len(arr)
	i, j := 0, n - 1
	for j - i >= k {
		if abs(arr[i] - x) <= abs(arr[j] - x) {
			j--
		} else {
			i++
		}
	}
	return arr[i:j+1]
}

func abs(a int) int {
	n := a >> 31
	return (n ^ a) - n
}

type node struct {
	val int
	diff int
}

func findClosestElements2(arr []int, k int, x int) []int {
	nodes := []node{}
	for _, v := range arr {
		n := node{ val: v }
		if v > x {
			n.diff = v - x
		} else {
			n.diff = x - v
		}
		nodes = append(nodes, n)
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].diff == nodes[j].diff {
			return nodes[i].val < nodes[j].val
		}
		return nodes[i].diff < nodes[j].diff
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = nodes[i].val
	}

	sort.Ints(res)

	return res
}
// @lc code=end

