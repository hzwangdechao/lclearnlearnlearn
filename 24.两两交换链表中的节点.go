/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (65.38%)
 * Likes:    543
 * Dislikes: 0
 * Total Accepted:    122K
 * Total Submissions: 184.6K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
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
func swapPairs(head *ListNode) *ListNode {
	// return recursive(head)
	return iteration(head)

}

func iteration(head *ListNode) *ListNode{
	temp := &ListNode{}
	temp.Next = head
	preNode := temp

	for head != nil && head.Next != nil{
		firstNode := head
		secondNode := head.Next

		preNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		fmt.Println(head)
		preNode = firstNode
		head = head.Next

	}

	return temp.Next
}


func recursive(head *ListNode)*ListNode  {
	// 空链表的情况或一个节点的情况直接返回
	if head == nil || head.Next == nil{
		return head
	}

	firstNode := head
	nextNode := head.Next

	firstNode.Next = recursive(nextNode.Next)
	nextNode.Next = firstNode

	return nextNode

}

// @lc code=end
