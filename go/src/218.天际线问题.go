/*
 * @lc app=leetcode.cn id=218 lang=golang
 *
 * [218] 天际线问题
 */

// @lc code=start
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type pos struct {
	x int
	height int
}


func getSkyline(buildings [][]int) [][]int {
	n, i := len(buildings), 0
	all := make([]*pos, n*2)
	for _, v := range buildings {
		all[i] = &pos{
			x: v[0],
			height: -v[2], // 起点高度用负数
		}
		all[i + 1] = &pos{
			x: v[1],
			height: v[2],
		}
		i += 2
	}

	// 根据x坐标从左到右排列所有点
	sort.Slice(all, func(i, j int) bool { 
		if all[i].x == all[j].x {
			return (all[i].height) < (all[j].height)
		}
		return all[i].x < all[j].x
	})

	h := &IntHeap{} // 大顶堆 存储高度
	heap.Init(h)
	res, last := [][]int{}, 0
	for i := 0; i < 2 * n; i++ {
		x, y := all[i].x, all[i].height

		if y > 0 { // 右边界
			// 从堆中移除该点的高度
			j := 0
			for ; j < h.Len(); j++ {
				if (*h)[j] == y {
					break
				}
			}
			heap.Remove(h, j)

			if h.Len() == 0 {
				res = append(res, []int{x, 0})
			} else if y > (*h)[0] { 
				if y == last { // 比剩下的元素高，剩下的元素不可能在该x坐标处结尾，因为前面对all排序时，当x相同时，y坐标是升序排序的
					res = append(res, []int{x, (*h)[0]})
					last = (*h)[0]
				} else { // 似乎不会执行到这里
					res = append(res, []int{x, y})
					last = y
				}
			}
		} else {
			if h.Len() == 0 || -y > (*h)[0] {
				res = append(res, []int{x, -y})
				last = -y
			}
			heap.Push(h, -y)
		}
	}

	return res
}

func abs(n int) int {
	y := n >> 31          // y ← x >> 31
	return (n ^ y) - y    // (x ⨁ y) - y
}
// @lc code=end

