/*
 * @lc app=leetcode.cn id=901 lang=rust
 *
 * [901] 股票价格跨度
 */

// @lc code=start
struct StockSpanner {
    // (price, index)
    stack: Vec<(i32, i32)>,
    cur_day: i32,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl StockSpanner {
    fn new() -> Self {
        Self {
            stack: vec![(i32::MAX, -1)],
            cur_day: -1,
        }
    }

    fn next(&mut self, price: i32) -> i32 {
        while self.stack.last().unwrap().0 <= price {
            self.stack.pop();
        }

        self.cur_day += 1;
        let res = self.cur_day - self.stack.last().unwrap().1;
        self.stack.push((price, self.cur_day));

        res
    }
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * let obj = StockSpanner::new();
 * let ret_1: i32 = obj.next(price);
 */
// @lc code=end

