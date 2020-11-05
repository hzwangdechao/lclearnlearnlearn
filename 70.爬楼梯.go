/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (48.47%)
 * Likes:    970
 * Dislikes: 0
 * Total Accepted:    185.1K
 * Total Submissions: 381.3K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

// @lc code=start
// func climbStairs(n int) int {
// 	// 动态规划
// 	if n ==1 {
// 		return 1
// 	}
// 	if n == 2{
// 		return 2
// 	}
// 	a, b := 1, 2
// 	for i := 3; i <n ;i++{
// 		a, b = b, a+b
// 	}
// 	return a + b
// }

// func climbStairs(n int) int {
// 	// 递归解法, 这样写会 超时
// 	if n == 0{
// 		return 1
// 	}
// 	if n < 0{
// 		return 0
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {

	if n < 3{
		return n
	}

	x := 1
	y := 2
	for i:=2; i<n; i++{
		x, y = y, x+y
	}
	return y



	// // 引入临时变量实现递归思想
	// if n == 1{
	// 	return 1
	// }
	// if n == 2{
	// 	return 2
	// }
	// var tmp int
	// x, y := 1, 2
	// for i := 2; i<n; i++{
	// 	tmp = x
	// 	x = y
	// 	y = tmp + y
	// }
	// return y
}


// @lc code=end
