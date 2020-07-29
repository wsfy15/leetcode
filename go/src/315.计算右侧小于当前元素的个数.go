/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start
// 二叉搜索树
type treenode struct {
	Left *treenode
	Right *treenode
	Val int
	Count int // 左子树节点个数
}

func countSmaller(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	res := make([]int, n)
	var root *treenode
	// 将比较过的节点加入二叉搜索树，新节点去二叉搜索树中找到比它小的节点的左子树节点数量
	for i := n - 1; i >= 0; i-- {
		root = addNode(root, nums[i], i, res)
	}

	return res
}

func addNode(root *treenode, val, index int, res []int) *treenode {
	if root == nil {
		return &treenode{
			Val: val,
			Count: 0,
		}
	}

	if root.Val < val {
		res[index] += root.Count + 1
		root.Right = addNode(root.Right, val, index, res)
	} else {
		root.Count++
		root.Left = addNode(root.Left, val, index, res)
	}

	return root
}
// @lc code=end

