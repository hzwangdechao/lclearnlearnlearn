/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (65.91%)
 * Likes:    101
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 9.5K
 * Testcase Example:  '[1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]'
 *
 * 返回与给定的前序和后序遍历匹配的任何二叉树。
 *
 * pre 和 post 遍历中的值是不同的正整数。
 *
 *
 *
 * 示例：
 *
 * 输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
 * 输出：[1,2,3,4,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= pre.length == post.length <= 30
 * pre[] 和 post[] 都是 1, 2, ..., pre.length 的排列
 * 每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个。
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(post) == 0{
		return nil
	}

	if len(pre) == 1{
		node := &TreeNode{}
		node.Val = pre[0]
		return node
	}

	node := &TreeNode{}
	node.Val = pre[0]

	l := len(pre)
	t := pre[1]

	leftIndex := 0
	for i, v := range post{
		if v == t{
			leftIndex = i
			break
		}
	}
	node.Left = constructFromPrePost(pre[1:leftIndex+2], post[0:leftIndex+1])
	node.Right = constructFromPrePost(pre[leftIndex+2:], post[leftIndex+1:l-1])

	return node
}
// @lc code=end
