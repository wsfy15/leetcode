/*
 * @lc app=leetcode.cn id=993 lang=rust
 *
 * [993] 二叉树的堂兄弟节点
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
    pub fn is_cousins(root: Option<Rc<RefCell<TreeNode>>>, x: i32, y: i32) -> bool {
        fn dfs(node: Option<Rc<RefCell<TreeNode>>>, depth: i32, parent: i32, val: i32) -> Vec<i32> {
            if node.is_none() {
                return vec![];
            }

            let cur = node.clone().unwrap().borrow().val;
            if cur == val {
                return vec![depth, parent];
            }

            let res = dfs(
                node.clone().unwrap().borrow().left.clone(),
                depth + 1,
                cur,
                val,
            );
            if res.len() == 2 {
                return res;
            }
            return dfs(
                node.clone().unwrap().borrow().right.clone(),
                depth + 1,
                cur,
                val,
            );
        }

        let parent = root.clone().unwrap().borrow().val;
        // xdp[0]: depth xdp[1]: parent
        let xdp = dfs(root.clone(), 0, parent, x);
        let ydp = dfs(root.clone(), 0, parent, y);
        xdp[0] == ydp[0] && xdp[1] != ydp[1]
    }
}
// @lc code=end
