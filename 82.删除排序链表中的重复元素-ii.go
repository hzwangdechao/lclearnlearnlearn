/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (47.46%)
 * Likes:    271
 * Dislikes: 0
 * Total Accepted:    45.6K
 * Total Submissions: 95.8K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
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
	// 空链表的情况
	if head == nil{
		return head
	}
	// 记录每个值的数量
	m := make(map[int]int)
	temp1 := head
	for {
		m[temp1.Val] += 1
		if temp1.Next == nil{
			break
		}
		temp1 = temp1.Next
	}

	temp := &ListNode{}
	temp.Next  = head
	pre, cur := temp, head
	for {
		cnt , _ := m[cur.Val]

		if cnt == 1{
			// 啥也不做， 指针向后移动
			pre = pre.Next
		}else{
			// 将数量大于1 的情况全部删除
			pre.Next = cur.Next
		}
		if cur.Next == nil{
			break
		}
		cur = cur.Next
	}

	return temp.Next

}
// @lc code=end
