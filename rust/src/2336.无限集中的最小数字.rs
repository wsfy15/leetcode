/*
 * @lc app=leetcode.cn id=2336 lang=rust
 *
 * [2336] 无限集中的最小数字
 */

// @lc code=start
struct SmallestInfiniteSet {
    min: i32,
    set: std::collections::BTreeSet<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl SmallestInfiniteSet {
    fn new() -> Self {
        Self {
            min: 1,
            set: std::collections::BTreeSet::new(),
        }
    }

    fn pop_smallest(&mut self) -> i32 {
        if let Some(v) = self.set.iter().next() {
            let x = v.clone();
            self.set.remove(&x);
            x
        } else {
            let res = self.min;
            self.min += 1;
            res
        }
    }

    fn add_back(&mut self, num: i32) {
        if num < self.min {
            self.set.insert(num);
        }
    }
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * let obj = SmallestInfiniteSet::new();
 * let ret_1: i32 = obj.pop_smallest();
 * obj.add_back(num);
 */
// @lc code=end

