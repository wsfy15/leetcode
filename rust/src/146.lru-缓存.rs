/*
 * @lc app=leetcode.cn id=146 lang=rust
 *
 * [146] LRU 缓存
 */

// @lc code=start
use std::cell::{Ref, RefCell};
use std::collections::{BinaryHeap, HashMap, HashSet};
use std::rc::Rc;

struct Node {
    key: i32,
    val: i32,
    prev: Option<Rc<RefCell<Node>>>,
    next: Option<Rc<RefCell<Node>>>,
}

impl Node {
    pub fn new(key: i32, value: i32) -> Rc<RefCell<Node>> {
        Rc::new(RefCell::new(Self {
            key: key,
            val: value,
            prev: None,
            next: None,
        }))
    }
}

struct LRUCache {
    capacity: usize,
    map: HashMap<i32, Rc<RefCell<Node>>>,
    dummy: Rc<RefCell<Node>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl LRUCache {
    pub fn new(capacity: i32) -> Self {
        let capacity = capacity as usize;
        let dummy = Node::new(0, 0);
        dummy.borrow_mut().next = Some(dummy.clone());
        dummy.borrow_mut().prev = Some(dummy.clone());
        let map = HashMap::new();
        Self {
            capacity,
            map,
            dummy,
        }
    }

    pub fn get(&mut self, key: i32) -> i32 {
        if let Some(node) = self.map.get(&key) {
            let value = node.borrow().val;
            let node = node.clone();
            self.remove(node.clone());
            self.push_front(node);

            return value;
        }

        -1
    }

    pub fn put(&mut self, key: i32, value: i32) {
        if let Some(node) = self.map.get(&key) {
            let node = node.clone();
            node.borrow_mut().val = value;
            self.remove(node.clone());
            self.push_front(node);
            return;
        }

        let node = Node::new(key, value);
        self.map.insert(key, node.clone());
        self.push_front(node);

        if self.map.len() > self.capacity {
            let last_node = self.dummy.borrow().prev.clone().unwrap();
            self.map.remove(&last_node.borrow().key);
            self.remove(last_node);
        }
    }

    fn remove(&mut self, node: Rc<RefCell<Node>>) {
        let prev = node.borrow().prev.clone().unwrap();
        let next = node.borrow().next.clone().unwrap();
        prev.borrow_mut().next = Some(next.clone());
        next.borrow_mut().prev = Some(prev.clone());
    }

    fn push_front(&mut self, node: Rc<RefCell<Node>>) {
        let next = self.dummy.borrow().next.clone();
        node.borrow_mut().prev = Some(self.dummy.clone());
        node.borrow_mut().next = next.clone();
        self.dummy.borrow_mut().next = Some(node.clone());
        next.unwrap().borrow_mut().prev = Some(node);
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * let obj = LRUCache::new(capacity);
 * let ret_1: i32 = obj.get(key);
 * obj.put(key, value);
 */
// @lc code=end

