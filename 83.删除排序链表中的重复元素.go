/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (50.20%)
 * Likes:    302
 * Dislikes: 0
 * Total Accepted:    98K
 * Total Submissions: 195K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * 示例 1:
 *
 * 输入: 1->1->2
 * 输出: 1->2
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil {
		if fast.Val != slow.Val{
			// slow = slow.Next
			// slow.Val = fast.Val

			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}

	slow.Next = nil
	return head
}
func deleteDuplicates1(head *ListNode) *ListNode {
	// 空链表的情况
	if head ==  nil{
		return head
	}

	m := make(map[int]bool)

	temp := &ListNode{}
	temp.Next = head
	pre, cur := temp, head


	for {
		// 第一次遍历的情况, 先从m中尝试获取， 如果获取不到， 则指针移动到下一个节点， 将值放到map中， 否则进行删除操作
		_, ok := m[cur.Val]
		if ok{
			// 找到的情况, 进行删除
			pre.Next  = cur.Next
		}else{
			// 找不到的情况
			m[cur.Val] = true
			pre = pre.Next
		}
		if cur.Next == nil{
			break
		}
		cur = cur.Next
	}


	return temp.Next

}
// @lc code=end
