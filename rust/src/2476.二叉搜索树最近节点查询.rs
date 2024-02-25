/*
 * @lc app=leetcode.cn id=2476 lang=rust
 *
 * [2476] 二叉搜索树最近节点查询
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
    pub fn closest_nodes(root: Option<Rc<RefCell<TreeNode>>>, queries: Vec<i32>) -> Vec<Vec<i32>> {
        let nums = Self::inorder_traversal(root.clone());
        let n = nums.len();

        let mut res = vec![];
        for query in queries {
            let index = nums.partition_point(|&x| x < query);
            let right = if index < n { nums[index] } else { -1 };
            let left = if index < n && nums[index] == query {
                nums[index]
            } else if index > 0 {
                nums[index - 1]
            } else {
                -1
            };
            res.push(vec![left, right]);
        }

        res
    }

    fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
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
