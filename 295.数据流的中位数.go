/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (48.49%)
 * Likes:    275
 * Dislikes: 0
 * Total Accepted:    23.7K
 * Total Submissions: 48.6K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
*/

// @lc code=start
type MedianFinder struct {
	minHeap *IntMinHeap
	maxHeap *IntMaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(IntMinHeap),
		maxHeap: new(IntMaxHeap),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 所有要添加的元素， 必须先添加到最大堆中， 然后再取出最大堆的堆顶元素放到最小堆中， 保证最小堆中的任意元素大于最大堆的任意元素
	heap.Push(this.maxHeap, num)
	heap.Push(this.minHeap, (heap.Pop(this.maxHeap)).(int))

	// 维护两个堆的大小差， 使得最大堆的大小总是大于等于最小堆
	if len(*this.maxHeap) < len(*this.minHeap) {
		heap.Push(this.maxHeap, (heap.Pop(this.minHeap)).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(*this.minHeap) == len(*this.maxHeap) {
		return (float64((*this.maxHeap)[0]) + float64((*this.minHeap)[0])) / float64(2)
	} else {
		return float64((*this.maxHeap)[0])
	}
}

// 定义最大堆和最小堆
type IntMaxHeap []int
type IntMinHeap []int

func (h IntMaxHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMaxHeap) Less(i int, j int) bool {
	return h[i] >= h[j]
}

func (h IntMinHeap) Less(i int, j int) bool {
	return h[i] <= h[j]
}

func (h IntMaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntMinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	length := len(*h)
	val := (*h)[length-1]
	*h = (*h)[:length-1]

	return val
}
func (h *IntMinHeap) Pop() interface{} {
	length := len(*h)
	val := (*h)[length-1]
	*h = (*h)[:length-1]

	return val
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
