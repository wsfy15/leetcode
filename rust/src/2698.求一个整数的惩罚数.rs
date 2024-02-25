/*
 * @lc app=leetcode.cn id=2698 lang=rust
 *
 * [2698] 求一个整数的惩罚数
 */

// @lc code=start
static mut INITIALIZED: bool = false;
static mut PRE_SUM: [i32; 1001] = [0; 1001];

fn init_once() {
    unsafe {
        if INITIALIZED {
            return;
        }

        INITIALIZED = true;
        for i in 1..1001 {
            let s = (i * i).to_string().bytes().collect();
            PRE_SUM[i] = PRE_SUM[i - 1]
                + if dfs(&s, 0, 0, i as i32) {
                    (i * i) as i32
                } else {
                    0
                };
        }
    }
}

fn dfs(s: &Vec<u8>, index: usize, sum: i32, target: i32) -> bool {
    if index == s.len() {
        return sum == target;
    }
    let mut num = 0;
    for i in index..s.len() {
        num = num * 10 + (s[i] & 0xf) as i32;
        if dfs(s, i + 1, num + sum, target) {
            return true;
        }
    }

    false
}

impl Solution {
    pub fn punishment_number(n: i32) -> i32 {
        unsafe {
            init_once();
            PRE_SUM[n as usize]
        }
    }
}
// @lc code=end
