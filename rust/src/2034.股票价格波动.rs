/*
 * @lc app=leetcode.cn id=2034 lang=rust
 *
 * [2034] 股票价格波动
 */

// @lc code=start
use std::collections::{BTreeMap, BinaryHeap, HashMap, HashSet};

struct StockPrice {
    max_timestamp: i32,
    ts_to_price: HashMap<i32, i32>,
    price_count: BTreeMap<i32, i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl StockPrice {
    fn new() -> Self {
        Self {
            max_timestamp: 0,
            ts_to_price: HashMap::new(),
            price_count: BTreeMap::new(),
        }
    }

    fn update(&mut self, timestamp: i32, price: i32) {
        self.max_timestamp = self.max_timestamp.max(timestamp);
        if let Some(old_price) = self.ts_to_price.get(&timestamp) {
            if let Some(cnt) = self.price_count.get(old_price) {
                if *cnt == 1 {
                    self.price_count.remove(old_price);
                } else {
                    self.price_count.insert(*old_price, *cnt - 1);
                }
            }
        }

        self.ts_to_price.insert(timestamp, price);
        *self.price_count.entry(price).or_insert(0) += 1;
    }

    fn current(&self) -> i32 {
        *self.ts_to_price.get(&self.max_timestamp).unwrap()
    }

    fn maximum(&self) -> i32 {
        *self.price_count.iter().rev().next().unwrap().0
    }

    fn minimum(&self) -> i32 {
        *self.price_count.iter().next().unwrap().0
    }
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * let obj = StockPrice::new();
 * obj.update(timestamp, price);
 * let ret_2: i32 = obj.current();
 * let ret_3: i32 = obj.maximum();
 * let ret_4: i32 = obj.minimum();
 */
// @lc code=end

