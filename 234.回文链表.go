/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (41.65%)
 * Likes:    497
 * Dislikes: 0
 * Total Accepted:    88.8K
 * Total Submissions: 212.4K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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


 // 递归的方法， 时间空间复杂度太大了
func methodWithRecursive(head *ListNode) bool {
	frontNode := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontNode.Val {
				return false
			}

			frontNode = frontNode.Next
		}
		return true
	}

	return recursivelyCheck(head)
}

func isPalindrome(head *ListNode) bool {
	return methodWithRecursive(head)
	if head == nil || head.Next == nil {
		return true
	}
	length := 0
	nodeArray := []int{}
	for head != nil {
		nodeArray = append(nodeArray, head.Val)
		length++
		head = head.Next
	}

	i, j := 0, length-1
	for i < j {
		if nodeArray[i] != nodeArray[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end
