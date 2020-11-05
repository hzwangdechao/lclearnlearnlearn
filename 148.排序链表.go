/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (65.16%)
 * Likes:    543
 * Dislikes: 0
 * Total Accepted:    60.9K
 * Total Submissions: 93.1K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 *
 * 示例 1:
 *
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 *
 *
 * 示例 2:
 *
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
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
func sortList(head *ListNode) *ListNode {
	// quickSort(head, nil)
	// return head

	return mergeSortList(head)
}

// 快速排序的方法对链表进行排序
func quickSort(head, end *ListNode) {
	if head == end || head.Next == end{
		return // 头节点或者头节点的下一个节点是结束节点， 停止遍历
	}
	temp := head.Val // 选取最左侧的元素作为基准元素
	slow, fast := head, head.Next
	for fast != nil{
		// 如果快指针小于基准元素的话， 将慢指针向后移动
		if fast.Val < temp{
			slow = slow.Next
			slow.Val, fast.Val = fast.Val, slow.Val
		}
		fast = fast.Next
	}
	// fast遍历最后时， slow指针指向的就是中间节点的位置
	head.Val, slow.Val = slow.Val, head.Val
	quickSort(head, slow)
	quickSort(slow.Next, end)
}


// 归并排序的方法对链表进行排序
func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	slow, fast := head, head
	// 使用快慢指针的方法， 将slow指针指向链表中间的位置
	for fast.Next != nil  && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	right := mergeSortList(slow.Next)
	slow.Next = nil
	left := mergeSortList(head)
	return mergeTwoSortedLists(left, right)
}


func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	res := temp
	for l1!=nil && l2!=nil{
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



// @lc code=end
