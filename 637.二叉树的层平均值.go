/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 *
 * https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (65.19%)
 * Likes:    174
 * Dislikes: 0
 * Total Accepted:    37.6K
 * Total Submissions: 55.2K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 输出：[3, 14.5, 11]
 * 解释：
 * 第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 节点值的范围在32位有符号整数范围内。
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
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil{
		return res
	}

	Q := []*TreeNode{root}
	nodeList := [][]int{}
	for i:=0; len(Q)>0; i++{
		nextQ := []*TreeNode{}
		nodeList = append(nodeList, []int{})
		for j:=0; j<len(Q); j++{
			curNode := Q[j]
			nodeList[i] = append(nodeList[i], curNode.Val)

			if curNode.Left != nil{
				nextQ = append(nextQ, curNode.Left)
			}

			if curNode.Right != nil{
				nextQ = append(nextQ, curNode.Right)
			}
		}

		sum := 0
		cnt := 0
		for _, num := range nodeList[i]{
			sum += num
			cnt += 1
		}

		res = append(res, float64(sum)/float64(cnt))

		Q = nextQ
	}

	return res
}
// @lc code=end
