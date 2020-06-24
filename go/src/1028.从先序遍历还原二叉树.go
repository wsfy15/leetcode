/*
 * @lc app=leetcode.cn id=1028 lang=golang
 *
 * [1028] 从先序遍历还原二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorder(S string) *TreeNode {
	n := len(S)
	i := getNumEndIndex(S, 0)
	num, _ := strconv.Atoi(S[:i])
	root := &TreeNode{ Val: num }
	stack := []*TreeNode{root}
	for i < n {
		l := getLevelEndIndex(S, i)
		level := l - i
		if len(stack) > level {
			stack = stack[:level]
		}

		n := getNumEndIndex(S, l)
		num, _ = strconv.Atoi(S[l:n])
		newNode := &TreeNode{ Val: num }
		if stack[len(stack) - 1].Left == nil {
			stack[len(stack) - 1].Left = newNode
		} else {
			stack[len(stack) - 1].Right = newNode
		}
		stack = append(stack, newNode)

		i = n
	}

	return root
}

func getNumEndIndex(S string, i int) int {
	j := i
	for j < len(S) && S[j] != '-' {
		j++
	}
	return j
}

func getLevelEndIndex(S string, i int) int {
	j := i
	for j < len(S) && S[j] == '-' {
		j++
	}
	return j
}
// @lc code=end

