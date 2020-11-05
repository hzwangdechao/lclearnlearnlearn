/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 *
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (64.49%)
 * Likes:    338
 * Dislikes: 0
 * Total Accepted:    66.7K
 * Total Submissions: 103.4K
 * Testcase Example:  '[6,2,8,0,4,7,9,null,null,3,5]\n2\n8'
 *
 * 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x
 * 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
 *
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
 * 输出: 6
 * 解释: 节点 2 和节点 8 的最近公共祖先是 6。
 *
 *
 * 示例 2:
 *
 * 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
 * 输出: 2
 * 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
 *
 *
 *
 * 说明:
 *
 *
 * 所有节点的值都是唯一的。
 * p、q 为不同节点且均存在于给定的二叉搜索树中。
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return methodWithIteration(root, p, q)
	// return methodWithRecursive(root, p, q)
}

func methodWithIteration(root, p, q *TreeNode) *TreeNode {
	for root != nil{
		if root.Val > p.Val && root.Val > q.Val{
			root = root.Left
		}else if root.Val < p.Val && root.Val < q.Val{
			root = root.Right
		}else{
			return root
		}
	}
	return nil
}

func methodWithRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil{
		return root
	}

	if root.Val > p.Val && root.Val > q.Val{
		return methodWithRecursive(root.Left, p, q)
	}else if root.Val < p.Val && root.Val < q.Val{
		return methodWithRecursive(root.Right, p, q)
	}else{
		return root
	}
}



func iteration(root, p, q *TreeNode)*TreeNode  {

	for root != nil{
		if root.Val > p.Val && root.Val > q.Val{
			root = root.Left
		}else if root.Val < p.Val && root.Val < q.Val{
			root = root.Right
		}else{
			return root
		}
	}
	return root
}

func recursive(root, p, q *TreeNode)*TreeNode  {
	if root == nil{
		return root
	}
	if root.Val > p.Val && root.Val > q.Val{
		return recursive(root.Left, p, q)
	}else if root.Val < p.Val && root.Val < q.Val{
		return recursive(root.Right, p, q)
	}else{
		return root
	}
}
// @lc code=end
