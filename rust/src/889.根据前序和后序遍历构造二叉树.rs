/*
 * @lc app=leetcode.cn id=889 lang=rust
 *
 * [889] 根据前序和后序遍历构造二叉树
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
    pub fn construct_from_pre_post(
        preorder: Vec<i32>,
        postorder: Vec<i32>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        let n = preorder.len();
        Self::dfs(&preorder, &postorder, 0, n - 1, 0, n - 1)
    }

    fn dfs(
        pre: &Vec<i32>,
        post: &Vec<i32>,
        prestart: usize,
        preend: usize,
        poststart: usize,
        postend: usize,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if prestart > preend {
            return None;
        }
        if prestart == preend {
            return Some(Rc::new(RefCell::new(TreeNode {
                val: pre[prestart],
                left: None,
                right: None,
            })));
        }
        let mut node = TreeNode {
            val: pre[prestart],
            left: None,
            right: None,
        };

        // 寻找左子树根节点
        let mut index = poststart;
        while post[index] != pre[prestart + 1] {
            index += 1;
        }

        node.left = Self::dfs(
            pre,
            post,
            prestart + 1,
            prestart + 1 + index - poststart,
            poststart,
            index,
        );
        node.right = Self::dfs(
            pre,
            post,
            prestart + 2 + index - poststart,
            preend,
            index + 1,
            postend,
        );

        Some(Rc::new(RefCell::new(node)))
    }
}
// @lc code=end
