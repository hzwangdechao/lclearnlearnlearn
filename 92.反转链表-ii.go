/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (51.10%)
 * Likes:    500
 * Dislikes: 0
 * Total Accepted:    73.3K
 * Total Submissions: 142.6K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
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

// 将链表分成 pre + rev + 剩余节点
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var pre *ListNode // <m的节点
	var rev *ListNode // 1<=m<=n需要反转的部分
	var pRev *ListNode // 反转链表的首个节点， 反转完成之后将其指向剩余节点部分

	cur := head
	idx := 1

	for cur != nil{
		if idx > n{
			// 剩余节点
			break
		}
		if idx >= m{
			// 需要反转的部分
			temp := cur.Next
			cur.Next = rev
			rev = cur
			if pRev == nil{
				pRev = rev
			}
			cur = temp
		}else{
			// 小于m的部分， 正常将指针向后移动
			pre = cur
			cur = cur.Next
		}
		idx += 1
	}

	pRev.Next = cur // 连接反转过后的节点和剩余节点


	if pre != nil{
		// 如果m前面有节点的话，则连接<m的节点和反转过后的节点
		pre.Next = rev
		return head
	}else{
		// 小于m没有节点， 直接返回反转过后的节点
		return rev
	}


}

// @lc code=end
