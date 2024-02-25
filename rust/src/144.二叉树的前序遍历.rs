/*
 * @lc app=leetcode.cn id=144 lang=rust
 *
 * [144] 二叉树的前序遍历
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
    pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        if root.is_none() {
            return vec![];
        }

        let mut res = vec![];
        let mut cur = root;
        let mut stack = vec![];
        while stack.len() > 0 || cur.is_some() {
            while cur.is_some() {
                let cc: Rc<RefCell<TreeNode>> = cur.clone().unwrap();
                res.push(cc.borrow().val);
                stack.push(cc);
                cur = cur.clone().unwrap().borrow().left.clone();
            }

            let node = stack.pop().unwrap();
            cur = node.borrow().right.clone();
        }

        res
    }
}
// @lc code=end
