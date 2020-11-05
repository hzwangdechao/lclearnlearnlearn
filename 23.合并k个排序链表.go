/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (52.69%)
 * Likes:    878
 * Dislikes: 0
 * Total Accepted:    162.4K
 * Total Submissions: 306.9K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// return linear(lists)
	// return partitionMerge(lists, 0, len(lists)-1)
	return priorityQueue(lists)
}
// 最小堆优先队列实现

func priorityQueue(lists []*ListNode) *ListNode {
	h := HeapMin{}
	heap.Init(&h)
	node := &ListNode{}
	cur := node
	for i:=0;i<len(lists); i++{
		if lists[i] == nil{
			continue
		}
		heap.Push(&h, lists[i])
	}

	for h.Len() > 0{
		min := heap.Pop(&h).(*ListNode)
		cur.Next = min
		cur = cur.Next
		if min.Next != nil{
			heap.Push(&h, min.Next)
		}
	}
	return node.Next
}


type HeapMin []*ListNode

func (h HeapMin) Len() int  {
	return len(h)
}

func (h HeapMin) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *HeapMin) Push(x interface{})  {
	*h = append(*h, x.(*ListNode))
}

func (h *HeapMin) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func (h *HeapMin) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// 分治法进行合并
func partitionMerge(lists []*ListNode, l, r int) *ListNode {
	if l == r{
		return lists[l]
	}
	if l > r{
		var x *ListNode
		return x
	}

	mid := (r - l)/2 + l

	return mergeTwoLists(partitionMerge(lists, l, mid), partitionMerge(lists, mid+1, r))

}



// 线性顺序方法
func linear(lists []*ListNode) *ListNode {
	var ans *ListNode
	ans = nil
	for _, node := range lists{
		ans = mergeTwoLists(ans, node)
	}
	return ans
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	temp := new(ListNode)
	res := temp
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			temp.Next = l1
			l1 = l1.Next
		}else{
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 != nil{
		temp.Next = l1
	}
	if l2 != nil{
		temp.Next = l2
	}

	return res.Next

}
// @lc code=end
