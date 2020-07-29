/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 */

// @lc code=start
type item struct {
	row int
	col int
	val int
}
type ItemHeap []*item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := &ItemHeap{}
	for i := 0; i < n; i++ {
		*h = append(*h, &item{
			row: i,
			col: 0,
			val: matrix[i][0],
		})
	}

	heap.Init(h)
	for i := 1; i < k; i++ {
		val := heap.Pop(h)
		it, _ := val.(*item)
		if it.col + 1 < n {
			heap.Push(h, &item{
				row: it.row,
				col: it.col + 1,
				val: matrix[it.row][it.col+1],
			})
		}
	}
	
	val := heap.Pop(h)
	it, _ := val.(*item)
	return it.val
}

// 二分法
func kthSmallest2(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n - 1][n - 1]

	for left < right {
		mid := left + (right - left) >> 1
		if countNotMoreThan(matrix, mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}

// 计数时不要一个一个比较，以行为单位找，找到每一行最后一个 <=target 的数即知道每一行有多少个数<=target
func countNotMoreThan(matrix [][]int, target int) int {
	n, count := len(matrix) - 1, 0
	x, y := n, 0
	for x >= 0 && y <= n {
		if matrix[x][y] <= target {
			count += x + 1
			y++
		} else {
			x--
		}
	}
	return count
}
// @lc code=end

