/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
type NumArray struct {
	tree []int
	n int
}


func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n * 2)
	for i := n; i < n * 2; i++ {
		tree[i] = nums[i - n]
	}
	for i := n - 1; i > 0; i-- {
		tree[i] = tree[i * 2] + tree[i * 2 + 1]
	}
	return NumArray{ tree: tree, n: n }
}


func (this *NumArray) Update(i int, val int)  {
	pos := this.n + i
	this.tree[pos] = val
	for pos > 0 {
		left, right := pos, pos
		if pos % 2 == 0 { // left node
			right++
		} else {
			left--
		}

		pos /= 2
		this.tree[pos] = this.tree[left] + this.tree[right]
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	i += this.n
	j += this.n
	for i <= j {
		if i & 1 == 1 { // 左节点是其父节点的右子，则加上该节点值，然后跳到右边表兄弟节点上
			sum += this.tree[i]
			i++
		}
		if j & 1 == 0 {
			sum += this.tree[j]
			j--
		}
		i /= 2
		j /= 2
	}

	return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
// @lc code=end

