/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (38.49%)
 * Likes:    816
 * Dislikes: 0
 * Total Accepted:    169.9K
 * Total Submissions: 440.4K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
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
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	// 两次遍历法， 先算出链表的长度， 要删除的节点就是 L-n+1

// 	// 第一次遍历求链表的长度
// 	length := 0
// 	temp1 := head
// 	for temp1 != nil{
// 		length += 1
// 		temp1  = temp1.Next
// 	}
// 	if length < 1 {
// 		return nil
// 	}
// 	temp := &ListNode{}
// 	temp.Next = head
// 	cur := temp
// 	index := length -n + 1  // 需要删除的节点的位置
// 	for cur != nil{
// 		if index == 1{
// 			cur.Next = cur.Next.Next
// 			break
// 		}
// 		index --
// 		cur = cur.Next
// 		fmt.Println(cur)
// 	}
// 	return temp.Next
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{}
	temp.Next = head

	first := temp  // 此链表 遍历到n+1个节点后从后向前遍历
	second := temp  // 此链表从开始向后遍历， 直到遍历n个节点()
	for i:=1 ; i <= n+1; i++{
		fmt.Println(i)
		first = first.Next
	}
	fmt.Println(first)
	for first != nil{
		first = first.Next
		second = second.Next
		fmt.Println(first)

	}
	second.Next = second.Next.Next
	return temp.Next


}
// @lc code=end
