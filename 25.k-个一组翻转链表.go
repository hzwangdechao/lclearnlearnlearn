/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (62.55%)
 * Likes:    710
 * Dislikes: 0
 * Total Accepted:    98.8K
 * Total Submissions: 157.2K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。
 *
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 *
 *
 * 示例：
 *
 * 给你这个链表：1->2->3->4->5
 *
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 *
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 *
 *
 *
 * 说明：
 *
 *
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil{
		return head
	}

	ahead := head
	for i:=1; i<k && ahead!=nil; i++{
		ahead = ahead.Next
	}
	if ahead == nil{
		// 说明遇上了不足k个节点的情况
		return head
	}

	// 将链表分成两个部分： 前面k个节点+后面的节点
	behind := ahead.Next  // 后面未翻转的链表
	ahead.Next = nil

	// 对前面的链表进行翻转
	newHead := reverseList(head)
	// 继续迭代后面的链表
	newBehind := reverseKGroup(behind, k)

	head.Next = newBehind

	return newHead
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
