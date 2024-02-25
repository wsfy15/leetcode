/*
 * @lc app=leetcode.cn id=106 lang=rust
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
    pub fn build_tree(
        mut inorder: Vec<i32>,
        mut postorder: Vec<i32>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if inorder.len() == 0 {
            return None;
        }

        let mut root = TreeNode {
            val: *postorder.last().unwrap(),
            left: None,
            right: None,
        };

        let index = inorder.iter().position(|&x| x == root.val).unwrap();

        let inorder_right = inorder.split_off(index);
        let postorder_right = postorder.split_off(index);

        root.left = Self::build_tree(inorder, postorder);
        root.right = Self::build_tree(
            inorder_right[1..].to_vec(),
            postorder_right[..postorder_right.len() - 1].to_vec(),
        );

        Some(Rc::new(RefCell::new(root)))
    }
}
// @lc code=end
