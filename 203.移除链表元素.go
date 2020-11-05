/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 *
 * https://leetcode-cn.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (45.27%)
 * Likes:    383
 * Dislikes: 0
 * Total Accepted:    77.6K
 * Total Submissions: 170.9K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * 删除链表中等于给定值 val 的所有节点。
 *
 * 示例:
 *
 * 输入: 1->2->6->3->4->5->6, val = 6
 * 输出: 1->2->3->4->5
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
func removeElements(head *ListNode, val int) *ListNode {

	temp := &ListNode{}
	temp.Next = head
	pre, cur := temp, head

	for cur != nil {
		// 遍历到最后一个节点
		if cur.Val != val {
			// 如果不想等的情况， 将pre移动到pre的Next节点
			pre = pre.Next
			// pre = cur
		} else {
			// 如果相等的情况， 将pre.Next 指向 cur.Next 这样就实现了删除的操作
			pre.Next = cur.Next
		}
		// 无论是否匹配， 都将当前节点向后移动一位
		cur = cur.Next

	}
	return temp.Next

}
// @lc code=end
