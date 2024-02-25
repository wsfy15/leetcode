/*
 * @lc app=leetcode.cn id=337 lang=rust
 *
 * [337] 打家劫舍 III
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
    pub fn rob(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::dfs(&root)[1]
    }

    // res[0] 代表不抢劫root的结果   res[1] 代表最优值
    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>) -> [i32; 2] {
        if let Some(root) = root {
            let left = Self::dfs(&root.borrow().left);
            let right = Self::dfs(&root.borrow().right);

            let val = left[0] + right[0] + root.borrow().val;
            let val2 = left[1] + right[1];
            [val2, val.max(val2)]
        } else {
            [0; 2]
        }
    }
}
// @lc code=end
