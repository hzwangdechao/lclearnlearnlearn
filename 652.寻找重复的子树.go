/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 *
 * https://leetcode-cn.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (53.54%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    9.6K
 * Total Submissions: 18K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 *
 * 两棵树重复是指它们具有相同的结构以及相同的结点值。
 *
 * 示例 1：
 *
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   2   4
 * ⁠      /
 * ⁠     4
 *
 *
 * 下面是两个重复的子树：
 *
 * ⁠     2
 * ⁠    /
 * ⁠   4
 *
 *
 * 和
 *
 * ⁠   4
 *
 *
 * 因此，你需要以列表的形式返回上述重复子树的根结点。
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	if root == nil{
		return res
	}
	hash := make(map[string]int, 0)
	helper(root, &res, hash)
	return res

}

func helper(root *TreeNode, res *[]*TreeNode, hash map[string]int) string {
	if root == nil{
		return ""
	}

	key := fmt.Sprintf("%s#%s#%s", root.Val, helper(root.Left, res, hash), helper(root.Right, res, hash))
	fmt.Println(key)
	if v, ok := hash[key]; ok{
		if v == 1{
			hash[key] ++
			*res = append(*res, root)
		}
	}else{
		hash[key] ++
	}

	return key
}
// @lc code=end
