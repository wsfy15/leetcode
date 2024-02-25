/*
 * @lc app=leetcode.cn id=1123 lang=rust
 *
 * [1123] 最深叶节点的最近公共祖先
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
    pub fn lca_deepest_leaves(
        root: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return None;
        }

        let root = root.unwrap();
        let left = Self::dfs(root.borrow().left.clone());
        let right = Self::dfs(root.borrow().right.clone());
        if left == right {
            return Some(root);
        }

        if left > right {
            return Self::lca_deepest_leaves(root.borrow().left.clone());
        }
        return Self::lca_deepest_leaves(root.borrow().right.clone());
    }

    fn dfs(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if root.is_none() {
            return 0;
        }
        let root = root.unwrap();
        let left = root.borrow().left.clone();
        let right = root.borrow().right.clone();
        1 + Self::dfs(left).max(Self::dfs(right))
    }
}
// @lc code=end
