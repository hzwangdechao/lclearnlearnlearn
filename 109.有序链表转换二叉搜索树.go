/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (72.96%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    34.9K
 * Total Submissions: 47.7K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 *
 * 示例:
 *
 * 给定的有序链表： [-10, -3, 0, 5, 9],
 *
 * 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
 *
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	 if head == nil{
		 return nil
	 }
	 var pre *ListNode
	 slow, fast := head, head

	 for fast != nil && fast.Next != nil{
		 fast = fast.Next.Next
		 pre = slow
		 slow = slow.Next
	 }
	//  if pre != nil{
	// 	 pre.Next = nil
	//  }


	 node := &TreeNode{
		 Val: slow.Val,
	 }
	 if slow == fast{
		 return node
	 }
	 node.Left = sortedListToBST(head)
	 node.Right = sortedListToBST(slow.Next)
	 return node
}

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}
// 快慢指针寻找中间位置的节点
func findMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}

	mid := findMedian(left, right)
	root := &TreeNode{mid.Val, nil, nil}

	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}

// @lc code=end
