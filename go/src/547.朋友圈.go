/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */

// @lc code=start
// 合并两个集合
type QuickUnionUF struct {
	Roots []int
	Count int
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	this := &QuickUnionUF{
		Roots: make([]int, n),
		Count: n,
	}
	for i := range this.Roots {
		this.Roots[i] = i
	}
	return this
}

// 返回元素i所属集合的root
func(this *QuickUnionUF) FindRoot(i int) int {
	root := i
	for this.Roots[root] != root { // 先找到最终的root
		root = this.Roots[root]
	}
	
	// 在这里进行优化，将从i到root的经过的节点的root指向最终的root
	for i != root {
		this.Roots[i], i = root, this.Roots[i]
	}
	
	return root
}

func(this *QuickUnionUF) Union(i, j int) {
	iroot := this.FindRoot(i)
	jroot := this.FindRoot(j)
	if iroot != jroot {
		this.Roots[iroot] = jroot
		this.Count--
	}
}

func findCircleNum(M [][]int) int {
	// 并查集
	n := len(M)
	if n <= 1 {
		return n
	}
	
	uf := NewQuickUnionUF(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Count
}
// @lc code=end

