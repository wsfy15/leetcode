/*
 * @lc app=leetcode.cn id=617 lang=rust
 *
 * [617] 合并二叉树
 */

// @lc code=start
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn merge_trees(root1: Option<Rc<RefCell<TreeNode>>>, root2: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        match (root1, root2) {
            (None, None) => return None,
            (some, None) | (None, some) => some,
            (Some(t1), Some(t2)) => match(t1.borrow_mut(), t2.borrow_mut()) {
                (mut t1, mut t2) => Some(Rc::new(RefCell::new(TreeNode{
                    val: t1.val + t2.val,
                    left: Self::merge_trees(t1.left.take(), t2.left.take()),
                    right: Self::merge_trees(t1.right.take(), t2.right.take()),
                })))
            } 
        }
    }
}
// @lc code=end

