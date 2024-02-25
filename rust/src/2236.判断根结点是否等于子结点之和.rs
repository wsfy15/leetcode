/*
 * @lc app=leetcode.cn id=2236 lang=rust
 *
 * [2236] 判断根结点是否等于子结点之和
 */

// @lc code=start
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }

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
    pub fn check_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(root) = root {
            let mut val = root.borrow().val;
            if let Some(left) = &root.borrow().left {
                val -= left.borrow().val;
            }
            if let Some(right) = &root.borrow().right {
                val -= right.borrow().val;
            }
            return val == 0;

        };
        return false;
    }
}
// @lc code=end

