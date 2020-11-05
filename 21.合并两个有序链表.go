/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (62.20%)
 * Likes:    1022
 * Dislikes: 0
 * Total Accepted:    258.1K
 * Total Submissions: 414.9K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
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
// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	preHead := &ListNode{}  // 假设一个空白链表

// 	result := preHead  // result.Next 就是最后需要的结果

// 	for l1 != nil && l2 != nil{
// 		// 如果l1的值小于l2的值
// 		if l1.Val < l2.Val {
// 			// 将头指针指向 l1
// 			preHead.Next = l1
// 			// l1 变为l1.Next
// 			l1 = l1.Next
// 		}else{
// 			// 将头指针指向l2
// 			preHead.Next = l2
// 			// l2变为l2.Next
// 			l2 = l2.Next
// 		}

// 		preHead = preHead.Next // l1或者l2
// 	}
// 	if l1 != nil{
// 		preHead.Next = l1
// 	}
// 	if l2 != nil{
// 		preHead.Next = l2
// 	}

// 	return result.Next
// }

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// 递归法
// 	if l1 == nil{
// 		return l2
// 	}
// 	if l2 == nil{
// 		return l1
// 	}

// 	if l1.Val <= l2.Val{
// 		l1.Next = mergeTwoLists(l1.Next, l2)
// 		return l1
// 	}else{
// 		l2.Next = mergeTwoLists(l1, l2.Next)
// 		return l2
// 	}

// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return recursive(l1, l2)

}


func iteration(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	res := temp

	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			temp.Next = l1
			l1 = l1.Next
		}else{
			temp.Next = l2
			l2 = l2.Next
		}

		temp = temp.Next
	}

	if l1 != nil{
		temp.Next = l1
	}
	if l2 != nil{
		temp.Next = l2
	}

	return res.Next
}

func recursive(l1 *ListNode, l2 *ListNode) *ListNode  {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	if l1.Val <= l2.Val{
		l1.Next = recursive(l1.Next, l2)
		return l1
	}else{
		l2.Next = recursive(l1, l2.Next)
		return l2
	}

}



// @lc code=end
