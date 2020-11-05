/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (68.78%)
 * Likes:    962
 * Dislikes: 0
 * Total Accepted:    238.5K
 * Total Submissions: 345.6K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
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
func reverseList(head *ListNode) *ListNode {
	return traverse(head)
}

func recursive(head *ListNode)*ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	res := recursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func traverse(head *ListNode)*ListNode {
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
