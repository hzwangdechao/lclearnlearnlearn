/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 *
 * https://leetcode-cn.com/problems/linked-list-cycle/description/
 *
 * algorithms
 * Easy (47.73%)
 * Likes:    594
 * Dislikes: 0
 * Total Accepted:    157.2K
 * Total Submissions: 328K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * 给定一个链表，判断链表中是否有环。
 *
 * 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：true
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 *
 *
 *
 *
 * 示例 2：
 *
 * 输入：head = [1,2], pos = 0
 * 输出：true
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 *
 *
 *
 *
 * 示例 3：
 *
 * 输入：head = [1], pos = -1
 * 输出：false
 * 解释：链表中没有环。
 *
 *
 *
 *
 *
 *
 * 进阶：
 *
 * 你能用 O(1)（即，常量）内存解决此问题吗？
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

 // hash 方法
// func hasCycle(head *ListNode) bool {

// 	m := make(map[*ListNode]bool)
// 	temp := head
// 	for temp != nil{
// 		_, ok := m[temp]
// 		if ok{
// 			return true
// 		}
// 		m[temp] = true
// 		temp = temp.Next

// 	}
// 	return false
// }

// 快慢指针方法
func hasCycle(head *ListNode) bool {
	// 空链表直接返回
	if head == nil{
		return false
	}
	fast := head.Next
	slow := head
	for fast != nil && slow != nil && fast.Next != nil{
		if fast == slow{
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
// @lc code=end
