/*
 * @lc app=leetcode.cn id=236 lang=rust
 *
 * [236] 二叉树的最近公共祖先
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
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn lowest_common_ancestor(
        root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return None;
        }

        if root.clone().unwrap().borrow().val == p.clone().unwrap().borrow().val
            || root.clone().unwrap().borrow().val == q.clone().unwrap().borrow().val
        {
            return root;
        }
        let left = Self::lowest_common_ancestor(
            root.clone().unwrap().borrow().left.clone(),
            p.clone(),
            q.clone(),
        );
        let right = Self::lowest_common_ancestor(
            root.clone().unwrap().borrow().right.clone(),
            p.clone(),
            q.clone(),
        );
        if left.is_some() && right.is_some() {
            return root;
        }
        if left.is_some() {
            return left;
        }
        return right;
    }
}
// @lc code=end
