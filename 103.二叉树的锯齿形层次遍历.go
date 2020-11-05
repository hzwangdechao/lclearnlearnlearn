/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (54.78%)
 * Likes:    261
 * Dislikes: 0
 * Total Accepted:    68.9K
 * Total Submissions: 125.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回锯齿形层次遍历如下：
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	res := [][]int{}

	Q := []*TreeNode{root}
	level := 1
	for i:=0;len(Q) > 0;i++{
		nextQ := make([]*TreeNode, 0)
		res = append(res, []int{})
		for j:=0; j<len(Q); j++{
			node := Q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil{
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil{
				nextQ = append(nextQ, node.Right)
			}
		}
		Q = nextQ
		if level % 2 == 0{
			reverseList(res[i])
		}
		level ++
	}
	return res
}

func reverseList(nums []int)  {
	l := 0
	r := len(nums)-1

	for l < r{
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}

}
// @lc code=end
