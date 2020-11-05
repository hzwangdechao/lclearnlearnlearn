/*
 * @lc app=leetcode.cn id=1028 lang=golang
 *
 * [1028] 从先序遍历还原二叉树
 *
 * https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/description/
 *
 * algorithms
 * Hard (73.83%)
 * Likes:    142
 * Dislikes: 0
 * Total Accepted:    17.2K
 * Total Submissions: 23.2K
 * Testcase Example:  '"1-2--3--4-5--6--7"'
 *
 * 我们从二叉树的根节点 root 开始进行深度优先搜索。
 *
 * 在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D +
 * 1。根节点的深度为 0）。
 *
 * 如果节点只有一个子节点，那么保证该子节点为左子节点。
 *
 * 给出遍历输出 S，还原树并返回其根节点 root。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入："1-2--3--4-5--6--7"
 * 输出：[1,2,5,3,4,6,7]
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入："1-2--3---4-5--6---7"
 * 输出：[1,2,5,3,null,6,null,4,null,7]
 *
 *
 * 示例 3：
 *
 *
 *
 * 输入："1-401--349---90--88"
 * 输出：[1,401,null,349,88,90]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 原始树中的节点数介于 1 和 1000 之间。
 * 每个节点的值介于 1 和 10 ^ 9 之间。
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
func recoverFromPreorder(S string) *TreeNode {
	return methodWithIteration(S)
	return methodWithMap(S)
}

func methodWithIteration(S string) *TreeNode {
	var stack []*TreeNode

	for i := 0; i < len(S); {
		var (
			depth int    = 0
			str   string = ""
		)

		for i < len(S) && S[i] == '-' {
			depth++
			i++
		}

		for i < len(S) && S[i] != '-' {
			str += string(S[i])
			i++
		}

		value, _ := strconv.Atoi(str)

		node := &TreeNode{Val: value}

		if len(stack) == 0 {
			stack = append(stack, node)
			continue
		}

		for len(stack) > depth {
			stack = stack[:len(stack)-1]
		}

		if stack[len(stack)-1].Left != nil {
			stack[len(stack)-1].Right = node
		} else {
			stack[len(stack)-1].Left = node
		}

		stack = append(stack, node)
	}

	return stack[0]
}

func methodWithMap(S string) *TreeNode {
	treeMap := make(map[int]*TreeNode)

	for i := 0; i < len(S); {
		var (
			depth int    = 0
			str   string = ""
		)

		// 一定要先判断- 再判断数字
		// - 的情况， 增加depth
		for i < len(S) && S[i] == '-' {
			depth++
			i++
		}

		// 数字的情况
		for i < len(S) && S[i] >= '0' && S[i] <= '9' {
			str += string(S[i])
			i++
		}

		value, _ := strconv.Atoi(str)
		node := &TreeNode{Val: value}

		if depth > 0 {
			if treeMap[depth-1].Left != nil {
				treeMap[depth-1].Right = node
			} else {
				treeMap[depth-1].Left = node
			}
		}
		treeMap[depth] = node
	}

	return treeMap[0]
}

// @lc code=end
