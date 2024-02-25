/*
 * @lc app=leetcode.cn id=1670 lang=rust
 *
 * [1670] 设计前中后队列
 */

// @lc code=start
// 两个双端队列，第二个长度始终大于等于第一个长度
use std::collections::VecDeque;

struct FrontMiddleBackQueue {
    first: VecDeque<i32>,
    second: VecDeque<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl FrontMiddleBackQueue {
    fn new() -> Self {
        Self {
            first: VecDeque::new(),
            second: VecDeque::new(),
        }
    }

    fn push_front(&mut self, val: i32) {
        self.first.push_front(val);
        if self.first.len() > self.second.len() {
            let val = self.first.pop_back().unwrap();
            self.second.push_front(val);
        }
    }

    fn push_middle(&mut self, val: i32) {
        if self.first.len() < self.second.len() {
            self.first.push_back(val)
        } else {
            self.second.push_front(val)
        }
    }

    fn push_back(&mut self, val: i32) {
        self.second.push_back(val);
        if self.second.len() > self.first.len() + 1 {
            let val = self.second.pop_front().unwrap();
            self.first.push_back(val);
        }
    }

    fn pop_front(&mut self) -> i32 {
        if self.first.is_empty() && self.second.is_empty() {
            return -1;
        }
        if self.first.is_empty() {
            return self.second.pop_front().unwrap();
        }

        let val = self.first.pop_front().unwrap();
        if self.second.len() > self.first.len() + 1 {
            let val = self.second.pop_front().unwrap();
            self.first.push_back(val);
        }
        val
    }

    fn pop_middle(&mut self) -> i32 {
        if self.first.is_empty() && self.second.is_empty() {
            return -1;
        }

        if self.first.is_empty() {
            return self.second.pop_front().unwrap();
        }

        if self.first.len() == self.second.len() {
            self.first.pop_back().unwrap()
        } else {
            self.second.pop_front().unwrap()
        }
    }

    fn pop_back(&mut self) -> i32 {
        if self.first.is_empty() && self.second.is_empty() {
            return -1;
        }
        let val = self.second.pop_back().unwrap();
        if self.first.len() > self.second.len() {
            let val = self.first.pop_back().unwrap();
            self.second.push_front(val);
        }
        val
    }
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * let obj = FrontMiddleBackQueue::new();
 * obj.push_front(val);
 * obj.push_middle(val);
 * obj.push_back(val);
 * let ret_4: i32 = obj.pop_front();
 * let ret_5: i32 = obj.pop_middle();
 * let ret_6: i32 = obj.pop_back();
 */
// @lc code=end

