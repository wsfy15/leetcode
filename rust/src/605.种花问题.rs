/*
 * @lc app=leetcode.cn id=605 lang=rust
 *
 * [605] 种花问题
 */

// @lc code=start
impl Solution {
    pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
        if n == 0 {
            return true;
        }
        let mut n = n;
        let mut prefix_ok = true;
        let mut index = 0;
        let len = flowerbed.len();
        while index < len {
            if prefix_ok {
                if flowerbed[index] == 0 && ((index == len - 1) || flowerbed[index + 1] == 0) {
                    n -= 1;
                    prefix_ok = false;
                    if n == 0 {
                        break;
                    }
                } else {
                    prefix_ok = flowerbed[index] == 0;
                }
            } else {
                prefix_ok = flowerbed[index] == 0;
            }
            index += 1;
        }
        n == 0
    }
}
// @lc code=end
