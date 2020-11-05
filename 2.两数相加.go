/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (37.30%)
 * Likes:    4280
 * Dislikes: 0
 * Total Accepted:    407.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建辅助节点
	temp := &ListNode{}
	head := temp
	sum := 0

	for l1 != nil || l2 != nil || sum > 0{
		temp.Next = &ListNode{}  // 初始化第一个节点

		temp = temp.Next

		if l1 != nil{
			sum = l1.Val + sum
			l1 = l1.Next
		}

		if l2 != nil{
			sum  = l2.Val + sum
			l2 = l2.Next
		}

		temp.Val = sum % 10  // 当前的位置取比10大的部分
		sum  = sum / 10
	}

	return head.Next

}


// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// 	res := &ListNode{} // 负责返回结果
// 	cur := res  // 负责移动指针
// 	sum := 0  // 两数相加之和
// 	carry := 0 // 负责进位

// 	for{
// 		sum = l1.Val + l2.Val + carry

// 		if sum / 10 > 0 {
// 			// 需要进位的情况
// 			carry = 1
// 		}else{
// 			// 不需要进位的情况
// 			carry = 0
// 		}
// 		sum  = sum % 10
// 		cur.Next = &ListNode{Val: sum}
// 		// 当前节点遍历完毕， 将指针移动向下一个节点
// 		cur  = cur.Next

// 		if l1.Next == nil && l2.Next == nil{
// 			// 两个链表都遍历完的情况
// 			if carry ==  1{
// 				// 如果还有进位的情况， 则向最后一位添加一个1
// 				cur.Next = &ListNode{Val: 1}
// 			}
// 			// 退出循环条件
// 			break
// 		}

// 		// 如果遍历完的情况， 则将值设置为0
// 		if l1.Next ==  nil{
// 			l1.Val = 0
// 		}else{
// 			l1 = l1.Next
// 		}

// 		if l2.Next == nil{
// 			l2.Val = 0
// 		}else{
// 			l2 = l2.Next
// 		}


// 	}
// 	return res.Next



// }

// @lc code=end
