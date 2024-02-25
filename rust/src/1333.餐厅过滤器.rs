/*
 * @lc app=leetcode.cn id=1333 lang=rust
 *
 * [1333] 餐厅过滤器
 */

// @lc code=start
impl Solution {
    pub fn filter_restaurants(
        restaurants: Vec<Vec<i32>>,
        vegan_friendly: i32,
        max_price: i32,
        max_distance: i32,
    ) -> Vec<i32> {
        let mut restaurants = restaurants
            .iter()
            .filter(|x| x[3] <= max_price && x[4] <= max_distance && x[2] >= vegan_friendly)
            .collect::<Vec<_>>();

        restaurants.sort_unstable_by_key(|x| (-x[1], -x[0]));
        restaurants.iter().map(|x| x[0]).collect()
    }
}
// @lc code=end
