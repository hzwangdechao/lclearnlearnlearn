/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (64.72%)
 * Likes:    324
 * Dislikes: 0
 * Total Accepted:    56.4K
 * Total Submissions: 86.2K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) []string {
	// return bfs(root)
	paths = []string{}
	dfs(root, "")
	return paths
}



// 广度遍历
func bfs(root *TreeNode) []string {
	paths := make([]string, 0)
	if root == nil{
		return paths
	}

	// 创建两个队列： 节点队列和路径队列
	nodeQueue := make([]*TreeNode, 0)
	pathQueue := make([]string, 0)

	// 初始化节点队列和路径队列
	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, strconv.Itoa(root.Val))

	// 遍历节点队列
	for i:=0; i<len(nodeQueue); i++{
		node := nodeQueue[i] // 当前节点
		path := pathQueue[i]  // 当前路径

		// 如果当前节点没有左右子树的话， 则将结果添加到paths中
		if node.Left == nil && node.Right == nil{
			paths = append(paths, path)
			continue
		}

		//  如果当前节点有左右子树的话，添加到队列中
		if node.Left != nil{
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path + "->" + strconv.Itoa(node.Left.Val))
		}

		if node.Right != nil{
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path + "->" + strconv.Itoa(node.Right.Val))
		}
	}

	return paths
}


// 深度遍历
var paths []string

func dfs(root *TreeNode, path string)  { 
	if root != nil{
		curPath := path
		curPath += strconv.Itoa(root.Val)

		if root.Left == nil && root.Right == nil{
			paths = append(paths, curPath)
		}else{
			dfs(root.Right, curPath+"->")
			dfs(root.Left, curPath+"->")
		}
	}

}
// @lc code=end
