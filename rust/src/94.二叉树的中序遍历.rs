/*
 * @lc app=leetcode.cn id=94 lang=rust
 *
 * [94] 二叉树的中序遍历
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
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        if root.is_none() {
            return vec![];
        }

        let mut res = vec![];
        let mut cur = root;
        let mut stack = vec![];
        while stack.len() > 0 || cur.is_some() {
            while cur.is_some() {
                stack.push(cur.clone().unwrap());
                cur = cur.clone().unwrap().borrow().left.clone();
            }

            let node = stack.pop().unwrap();
            res.push(node.borrow().val);
            cur = node.borrow().right.clone();
        }

        res
    }
}
// @lc code=end
