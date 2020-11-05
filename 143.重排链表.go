/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode-cn.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (56.34%)
 * Likes:    294
 * Dislikes: 0
 * Total Accepted:    36.8K
 * Total Submissions: 65.2K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
 * 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 * 示例 1:
 *
 * 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
 *
 * 示例 2:
 *
 * 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
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
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil{
		return
	}
	// 先使用快慢指针的方法， 将链表的后半部分进行一下翻转
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	// slow.Next 为后半部分
	right := slow.Next

	slow.Next = nil

	right = reverseList(right)

	// fmt.Println(right)
	left := head
	for right != nil{
		lNext := left.Next
		rNext := right.Next
		left.Next = right
		right.Next = lNext
		left = lNext
		right = rNext
	}

	// return res.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
// @lc code=end
