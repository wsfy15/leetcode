/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

 type Codec struct {
  sb strings.Builder
  index int 
}

func Constructor() Codec {
  return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
  if root == nil {
    return ""
  }
  
  this.sb = strings.Builder{}
  this.dfsSerialize(root)
  str := this.sb.String()
  return str[:len(str) - 1]
}

func (this *Codec) dfsSerialize(root *TreeNode) {
  if root == nil {
    this.sb.Write([]byte("n,"))
  } else {
    this.sb.Write([]byte(strconv.Itoa(root.Val) + ","))
    this.dfsSerialize(root.Left)
    this.dfsSerialize(root.Right)
  }
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
  if len(data) == 0 {
    return nil
  }
 
  strs := strings.Split(data, ",")
  this.index = 0
  return this.dfsDeserialize(strs)
}

func (this *Codec) dfsDeserialize(strs []string) *TreeNode {
  if this.index >= len(strs) || strs[this.index] == "n" {
    this.index++
    return nil
  }
  
  val, _ := strconv.Atoi(strs[this.index])
  root := &TreeNode{Val: val}
  this.index++
  root.Left = this.dfsDeserialize(strs)
  root.Right = this.dfsDeserialize(strs)
  
  return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */