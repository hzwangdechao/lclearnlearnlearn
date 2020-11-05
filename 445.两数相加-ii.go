··							/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (57.23%)
 * Likes:    204
 * Dislikes: 0
 * Total Accepted:    39.2K
 * Total Submissions: 68.5K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 *
 *
 * 进阶：
 *
 * 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
 *
 *
 *
 * 示例：
 *
 * 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 8 -> 0 -> 7
 *7243
 * 564
 *7807
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return withStack(l1, l2)
}

// 使用栈的方法
func withStack(l1 *ListNode, l2 *ListNode) *ListNode {
	// 使用两个栈， 将链表中的元素存储起来
	s1, s2 := []int{}, []int{}
	// 将元素添加到栈中
	for l1 != nil{
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil{
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	res := &ListNode{} // 返回最终的结果
	temp := res  // 临时变量， 用于链表的移动
	add := 0  // 进位数

	for len(s1)>0 || len(s2)>0 || add>0{
		// 当前数字的和
		sum := 0

		if len(s1) > 0{
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) > 0{
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		sum += add

		// 当前节点
		curNode := &ListNode{
			Val: sum%10,
		}
		curNode.Next = temp.Next
		temp.Next = curNode

		add = sum / 10
	}
	return res.Next
}

// @lc code=end
