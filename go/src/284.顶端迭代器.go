/*
 * @lc app=leetcode.cn id=284 lang=golang
 *
 * [284] 顶端迭代器
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	nextValid bool
	nextElement int
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
  return &PeekingIterator{ iter: iter }
}

func (this *PeekingIterator) hasNext() bool {
  return this.nextValid || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
  if this.nextValid {
		this.nextValid = false
		return this.nextElement
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
  if !this.nextValid {
		this.nextValid = true
		this.nextElement = this.iter.next()
  }
	return this.nextElement
}
// @lc code=end

