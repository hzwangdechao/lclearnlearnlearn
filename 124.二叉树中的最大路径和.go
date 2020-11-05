/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (42.96%)
 * Likes:    692
 * Dislikes: 0
 * Total Accepted:    73.2K
 * Total Submissions: 170K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空二叉树，返回其最大路径和。
 *
 * 本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[1,2,3]
 *
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 *
 * 输出：6
 *
 *
 * 示例 2：
 *
 * 输入：[-10,9,20,null,null,15,7]
 *
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 *
 * 输出：42
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
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode)int
	maxGain = func (node *TreeNode) int {
		// 如果节点为空的话， 最大贡献值为0
		if node == nil{
			return 0
		}

		leftGain := max(maxGain(node.Left), 0) // 左节点的最大贡献值
		rightGain := max(maxGain(node.Right), 0) // 右节点的最大贡献值

		newPricePath := leftGain + rightGain + node.Val

		maxSum = max(maxSum,newPricePath)

		return max(leftGain, rightGain) + node.Val
	}
	maxGain(root)
	return maxSum

}

func max(x int, y int) int {
	if x > y{
		return x
	}
	return y
}
// @lc code=end
