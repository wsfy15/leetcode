/*
 * @lc app=leetcode.cn id=1993 lang=rust
 *
 * [1993] 树上的操作
 */

// @lc code=start
struct LockingTree {
    parent: Vec<usize>,
    locked_by: Vec<Option<i32>>,
    children: Vec<Vec<usize>>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl LockingTree {
    fn new(parent: Vec<i32>) -> Self {
        let parent: Vec<_> = parent.into_iter().map(|x| x as usize).collect();
        let n = parent.len();
        let locked_by = vec![None; n];
        let mut children = vec![vec![]; n];
        for i in 1..n {
            children[parent[i]].push(i);
        }

        Self {
            parent,
            locked_by,
            children,
        }
    }

    fn lock(&mut self, num: i32, user: i32) -> bool {
        if let None = self.locked_by[num as usize] {
            self.locked_by[num as usize] = Some(user);
            true
        } else {
            false
        }
    }

    fn unlock(&mut self, num: i32, user: i32) -> bool {
        let num = num as usize;
        match self.locked_by[num] {
            Some(u) if u == user => {
                self.locked_by[num] = None;
                true
            }
            _ => false,
        }
    }

    fn upgrade(&mut self, num: i32, user: i32) -> bool {
        let num = num as usize;
        // 1. check not locked from current to root
        let mut p = num;
        while p != usize::MAX {
            if self.locked_by[p].is_some() {
                return false;
            }
            p = self.parent[p];
        }

        // 2. check children locked
        let mut success = false;
        let mut cd = self.children[num].clone();
        while let Some(node) = cd.pop() {
            if self.locked_by[node].is_some() {
                self.locked_by[node] = None;
                success = true;
            }

            cd.extend(self.children[node].iter());
        }

        if success {
            self.locked_by[num] = Some(user);
        }
        success
    }
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * let obj = LockingTree::new(parent);
 * let ret_1: bool = obj.lock(num, user);
 * let ret_2: bool = obj.unlock(num, user);
 * let ret_3: bool = obj.upgrade(num, user);
 */
// @lc code=end

