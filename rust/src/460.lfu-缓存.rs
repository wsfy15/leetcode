/*
 * @lc app=leetcode.cn id=460 lang=rust
 *
 * [460] LFU 缓存
 */

// @lc code=start
use std::cell::{Ref, RefCell};
use std::collections::{BinaryHeap, HashMap, HashSet};
use std::rc::Rc;

struct Node {
    key: i32,
    val: i32,
    freq: i32,
    prev: Option<Rc<RefCell<Node>>>,
    next: Option<Rc<RefCell<Node>>>,
}

impl Node {
    pub fn new(key: i32, value: i32) -> Rc<RefCell<Node>> {
        Rc::new(RefCell::new(Self {
            key: key,
            val: value,
            freq: 1,
            prev: None,
            next: None,
        }))
    }
}

struct LFUCache {
    capacity: usize,
    min_freq: i32,
    // key to node
    map: HashMap<i32, Rc<RefCell<Node>>>,
    // freq to node list
    freq: HashMap<i32, Rc<RefCell<Node>>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl LFUCache {
    pub fn new(capacity: i32) -> Self {
        Self {
            capacity: capacity as usize,
            min_freq: 0,
            map: HashMap::new(),
            freq: HashMap::new(),
        }
    }

    pub fn get(&mut self, key: i32) -> i32 {
        if let Some(node) = self.get_node(key) {
            return node.borrow().val;
        }

        -1
    }

    pub fn put(&mut self, key: i32, value: i32) {
        if let Some(node) = self.get_node(key) {
            node.borrow_mut().val = value;
            return;
        }

        if self.map.len() == self.capacity {
            let dummy = self.freq.get(&self.min_freq).unwrap();
            let dummy2 = &dummy.clone();
            let last_node = dummy.borrow().prev.clone().unwrap();
            let key = last_node.borrow().key;
            self.map.remove(&key);
            self.remove(last_node);
            if Rc::ptr_eq(dummy2, dummy2.borrow().prev.as_ref().unwrap()) {
                self.freq.remove(&self.min_freq);
            }
        }
        let node = Node::new(key, value);
        self.map.insert(key, node.clone());
        self.push_front(1, node.clone());
        self.min_freq = 1;
    }

    fn get_node(&mut self, key: i32) -> Option<Rc<RefCell<Node>>> {
        if let Some(node) = self.map.get(&key) {
            let node = node.clone();
            self.remove(node.clone());

            let freq = node.borrow().freq;
            let dummy = self.freq.get(&freq).unwrap();
            if Rc::ptr_eq(dummy, dummy.borrow().prev.as_ref().unwrap()) {
                self.freq.remove(&freq);
                if self.min_freq == freq {
                    self.min_freq += 1;
                }
            }

            node.borrow_mut().freq += 1;
            self.push_front(freq + 1, node.clone());
            return Some(node);
        }

        None
    }

    fn remove(&mut self, node: Rc<RefCell<Node>>) {
        let prev = node.borrow().prev.clone().unwrap();
        let next = node.borrow().next.clone().unwrap();
        prev.borrow_mut().next = Some(next.clone());
        next.borrow_mut().prev = Some(prev.clone());
    }

    fn push_front(&mut self, freq: i32, node: Rc<RefCell<Node>>) {
        let list = self.freq.entry(freq).or_insert_with(|| Self::new_list());
        let next = list.borrow().next.clone();
        node.borrow_mut().prev = Some(list.clone());
        node.borrow_mut().next = next.clone();
        list.borrow_mut().next = Some(node.clone());
        next.unwrap().borrow_mut().prev = Some(node);
    }

    fn new_list() -> Rc<RefCell<Node>> {
        let dummy = Node::new(0, 0);
        dummy.borrow_mut().prev = Some(dummy.clone());
        dummy.borrow_mut().next = Some(dummy.clone());
        dummy
    }
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * let obj = LFUCache::new(capacity);
 * let ret_1: i32 = obj.get(key);
 * obj.put(key, value);
 */
// @lc code=end

