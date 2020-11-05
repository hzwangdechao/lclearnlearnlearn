/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (44.49%)
 * Likes:    136
 * Dislikes: 0
 * Total Accepted:    23.2K
 * Total Submissions: 52K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n' +
  '[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * 设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。
 *
 * 你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用
 * KthLargest.add，返回当前数据流中第K大的元素。
 *
 * 示例:
 *
 *
 * int k = 3;
 * int[] arr = [4,5,8,2];
 * KthLargest kthLargest = new KthLargest(3, arr);
 * kthLargest.add(3);   // returns 4
 * kthLargest.add(5);   // returns 5
 * kthLargest.add(10);  // returns 5
 * kthLargest.add(9);   // returns 8
 * kthLargest.add(4);   // returns 8
 *
 *
 * 说明:
 * 你可以假设 nums 的长度≥ k-1 且k ≥ 1。
 *
 */

// @lc code=start

type IntHead []int

func (In *IntHead) Len() int {
	return len(*In)
}

func (In *IntHead) Less(i, j int) bool {
	return (*In)[i] < (*In)[j]
}
func (In *IntHead) Pop() (v interface{}) {
	*In, v = (*In)[:len(*In)-1], (*In)[len(*In)-1]
	return
}
func (In *IntHead) Push(x interface{}) {
	*In = append(*In, x.(int))
}
func (In *IntHead) Swap(i, j int) {
	(*In)[i], (*In)[j] = (*In)[j], (*In)[i]
}

type KthLargest struct {
	k int
	H *IntHead
}


func Constructor(k int, nums []int) KthLargest {
	min := &IntHead{}

	heap.Init(min)
	kth := KthLargest{
		k: k,
		H: min,
	}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth

}


func (this *KthLargest) Add(val int) int {
	//如果长度小于k，直接push
	if this.H.Len() < this.k {
		heap.Push(this.H, val)
	//判断val是否大于堆顶元素
	}else if (*this.H)[0] < val {
		heap.Pop(this.H)
		heap.Push(this.H, val)
	}
	return (*this.H)[0]

}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
