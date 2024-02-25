/*
 * @lc app=leetcode.cn id=307 lang=rust
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
struct NumArray {
    tree: Vec<i32>,
    n: usize,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        let n = nums.len();
        let mut tree = vec![0; 2 * n];
        for i in n..n * 2 {
            tree[i] = nums[i - n];
        }
        for i in (1..n).rev() {
            tree[i] = tree[i * 2] + tree[i * 2 + 1];
        }
        Self { tree, n }
    }

    fn update(&mut self, index: i32, val: i32) {
        let mut index = index as usize + self.n;
        self.tree[index] = val;
        while index > 0 {
            let (mut left, mut right) = (index, index);
            if index % 2 == 0 {
                //left node
                right += 1;
            } else {
                left -= 1;
            }
            index /= 2;
            self.tree[index] = self.tree[left] + self.tree[right];
        }
    }

    fn sum_range(&self, left: i32, right: i32) -> i32 {
        let mut sum = 0;
        let (mut left, mut right) = (left as usize + self.n, right as usize + self.n);
        while left <= right {
            if left & 1 == 1 {
                // 左节点是其父节点的右子，则加上该节点值，然后跳到右边表兄弟节点上
                sum += self.tree[left];
                left += 1;
            }
            if right & 1 == 0 {
                sum += self.tree[right];
                right -= 1;
            }

            left /= 2;
            right /= 2;
        }
        sum
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * let obj = NumArray::new(nums);
 * obj.update(index, val);
 * let ret_2: i32 = obj.sum_range(left, right);
 */
// @lc code=end

