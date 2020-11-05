/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 *
 * https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (45.95%)
 * Likes:    197
 * Dislikes: 0
 * Total Accepted:    32.4K
 * Total Submissions: 66.4K
 * Testcase Example:  '[1,null,2,2]'
 *
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 *
 * 假定 BST 有如下定义：
 *
 *
 * 结点左子树中所含结点的值小于等于当前结点的值
 * 结点右子树中所含结点的值大于等于当前结点的值
 * 左子树和右子树都是二叉搜索树
 *
 *
 * 例如：
 * 给定 BST [1,null,2,2],
 *
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 *
 *
 * 返回[2].
 *
 * 提示：如果众数超过1个，不需考虑输出顺序
 *
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
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

// 重复出现的数字一定是连续出现的
// 用base记录当前的数字
// count 记录当前数字重复的次数
// maxCount 已经扫描过的数字中出现最多的那个数字的出现次数
// answer 记录出现的众数

// 首先更新base和count    如果当前元素和base相等，count+=1   如果和当前元素不相等的话， 更新base为当前元素， 并且count变成1
// 然后更新maxCount 如果count==maxCount，说明这个base数字出现的次数等于当前众数出现的次数， 将base加入到answer数组中。   如果count>maxCount， 说明这个base数字出现的次数大于当前众数出现的次数， 因此需要将maxCount更新为count， 清空answer后将base添加到answer中

// func findMode(root *TreeNode) (answer []int) {
//     var base, count, maxCount int
//     update := func(x int) {
//         if x == base {
//             count++
//         } else {
//             base, count = x, 1
//         }
//         if count == maxCount {
//             answer = append(answer, base)
//         } else if count > maxCount {
//             maxCount = count
//             answer = []int{base}
//         }
//     }
//     cur := root
//     for cur != nil {
//         if cur.Left == nil {
//             update(cur.Val)
// 			cur = cur.Right
//             continue
// 		}

//         pre := cur.Left
//         for pre.Right != nil && pre.Right != cur {
//             pre = pre.Right
//         }
//         if pre.Right == nil {
//             pre.Right = cur
//             cur = cur.Left
//         } else {
//             pre.Right = nil
//             update(cur.Val)
//             cur = cur.Right
//         }
//     }
//     return
// }

// 递归的方式
func findModeWithDfs(root *TreeNode) (answer []int) {
	// 递归方式  左  根   右
	var base, count, maxCount int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		nodeVal := node.Val
		if nodeVal == base {
			count++
		} else {
			base = nodeVal
			count = 1
		}

		if count == maxCount {
			answer = append(answer, nodeVal)
		} else if count > maxCount {
			answer = []int{nodeVal}
			maxCount = count
		}

		dfs(node.Right)
	}

	dfs(root)
	return
}



// @lc code=end
