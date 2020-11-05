/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 *
 * https://leetcode-cn.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (40.53%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    83.4K
 * Total Submissions: 205.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->4->5->NULL, k = 2
 * 输出: 4->5->1->2->3->NULL
 * 解释:
 * 向右旋转 1 步: 5->1->2->3->4->NULL
 * 向右旋转 2 步: 4->5->1->2->3->NULL
 *
 *
 * 示例 2:
 *
 * 输入: 0->1->2->NULL, k = 4
 * 输出: 2->0->1->NULL
 * 解释:
 * 向右旋转 1 步: 2->0->1->NULL
 * 向右旋转 2 步: 1->2->0->NULL
 * 向右旋转 3 步: 0->1->2->NULL
 * 向右旋转 4 步: 2->0->1->NULL
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <=0{
		return head
	}
	// 首先将链表捣鼓成一个环形链表, 并且获取链表的场地
	l := 1
	cur := head
	for cur.Next != nil{
		l += 1
		cur = cur.Next
	}
	cur.Next = head  // 链表成员

	for i:=0; i<(l-k%l); i++{
		// 到达新链表的尾部节点
		cur = cur.Next
	}

	// 新的尾部与头部节点断开连接， 返回新的头部节点
	head, cur.Next = cur.Next, nil
    return head

}
// @lc code=end
