/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
type linkedList struct {
	val  []int
	next *linkedList
}

func (l *linkedList) add(index int, val []int) {
	cur := l
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = &linkedList{
		val:  val,
		next: cur.next,
	}
}

func (l linkedList) toSlice() [][]int {
	res := [][]int{}
	cur := l.next
	for cur != nil {
		res = append(res, cur.val)
		cur = cur.next
	}
	return res
}

func reconstructQueue(people [][]int) [][]int {
	// 先根据身高降序、k升序排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	// 根据每个people的k插入到指定位置
	list := linkedList{next: nil}
	for _, v := range people {
		list.add(v[1], v)
	}

	return list.toSlice()
}

// @lc code=end

