/*
 * @lc app=leetcode.cn id=105 lang=rust
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
        mut preorder: Vec<i32>,
        mut inorder: Vec<i32>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if preorder.len() == 0 {
            return None;
        }

        let mut root = TreeNode {
            val: preorder[0],
            left: None,
            right: None,
        };

        let index = inorder.iter().position(|&x| x == preorder[0]).unwrap();

        let preorder_right = preorder.split_off(index + 1);
        let inorder_right = inorder.split_off(index);

        root.left = Self::build_tree(preorder[1..].to_vec(), inorder);
        root.right = Self::build_tree(preorder_right, inorder_right[1..].to_vec());

        Some(Rc::new(RefCell::new(root)))
    }
}
// @lc code=end
