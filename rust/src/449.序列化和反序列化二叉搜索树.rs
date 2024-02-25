/*
 * @lc app=leetcode.cn id=449 lang=rust
 *
 * [449] 序列化和反序列化二叉搜索树
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
use std::rc::Rc;
use std::cell::RefCell;
struct Codec {
	
}

/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Codec {
    fn new() -> Self {
        Codec{}
    }

    fn serialize(&self, root: Option<Rc<RefCell<TreeNode>>>) -> String {
        if root.is_none() {
            return "".to_string();
        }

        let mut str = String::new();

        let mut queue = vec![root];
        while queue.len() > 0 {
            let node = queue.remove(0);
            if let Some(node) = node {
                str.push_str(&node.borrow().val.to_string());
                queue.push(node.borrow_mut().left.clone());
                queue.push(node.borrow_mut().right.clone());
            }
            str.push_str(",");
        }

        str
    }
	
    fn deserialize(&self, data: String) -> Option<Rc<RefCell<TreeNode>>> {
        if data == "" {
            return None;
        }
    
        let items: Vec<_> = data.split(",").collect();
    
        let val = items[0].parse().unwrap();
        let root = Rc::new(RefCell::new(TreeNode::new(val)));
    
        let mut queue = vec![Some(root.clone())];
        let mut index = 1;
        while queue.len() > 0 {
            let node = queue.remove(0);
            if let Some(node) = node {
                if items[index].len() > 0 {
                    let val = items[index].parse().unwrap();
                    let new_node = Rc::new(RefCell::new(TreeNode::new(val)));
                    node.borrow_mut().left = Some(new_node.clone());
                    queue.push(Some(new_node));
                }
                if items[index + 1].len() > 0 {
                    let val = items[index + 1].parse().unwrap();
                    let new_node = Rc::new(RefCell::new(TreeNode::new(val)));
                    node.borrow_mut().right = Some(new_node.clone());
                    queue.push(Some(new_node));
                }
                index += 2;
            }
        }
    
        Some(root)
    }
}

/**
 * Your Codec object will be instantiated and called as such:
 * let obj = Codec::new();
 * let data: String = obj.serialize(strs);
 * let ans: Option<Rc<RefCell<TreeNode>>> = obj.deserialize(data);
 */
// @lc code=end

