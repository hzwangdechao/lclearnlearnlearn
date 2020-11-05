/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 *
 * https://leetcode-cn.com/problems/linked-list-cycle-ii/description/
 *
 * algorithms
 * Medium (50.05%)
 * Likes:    466
 * Dislikes: 0
 * Total Accepted:    76.7K
 * Total Submissions: 152.6K
 * Testcase Example:  '[3,2,0,-4]\n1'
 *
 * 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
 *
 * 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
 *
 * 说明：不允许修改给定的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：tail connects to node index 1
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 *
 *
 *
 *
 * 示例 2：
 *
 * 输入：head = [1,2], pos = 0
 * 输出：tail connects to node index 0
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 *
 *
 *
 *
 * 示例 3：
 *
 * 输入：head = [1], pos = -1
 * 输出：no cycle
 * 解释：链表中没有环。
 *
 *
 *
 *
 *
 *
 * 进阶：
 * 你是否可以不用额外空间解决此题？
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

 // 与141题基本相似， 141题是成环链表时返回true， 142题是返回链表


// func detectCycle(head *ListNode) *ListNode {
// 	if head == nil{
// 		return head
// 	}
// 	m := make(map[*ListNode]bool)
// 	temp := head
// 	for temp != nil{
// 		if _, ok := m[temp]; ok{
// 			return temp
// 		}
// 		m[temp] = true
// 		temp = temp.Next
// 	}
// 	return nil

// }

// 快慢指针法
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}
	fast , slow := head, head
	// 快指针一次走两次， 慢指针一次走一次
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast{
			// 快指针和慢指针相遇了

			// 将满指针指向head， 快指针的步伐变为和慢指针一样
			slow = head
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil

}

// @lc code=end
