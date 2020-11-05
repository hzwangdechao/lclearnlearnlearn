/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 *
 * https://leetcode-cn.com/problems/4sum-ii/description/
 *
 * algorithms
 * Medium (55.32%)
 * Likes:    179
 * Dislikes: 0
 * Total Accepted:    24.7K
 * Total Submissions: 44.5K
 * Testcase Example:  '[1,2]\n[-2,-1]\n[-1,2]\n[0,2]'
 *
 * 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] +
 * D[l] = 0。
 *
 * 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28 - 1
 * 之间，最终结果不会超过 2^31 - 1 。
 *
 * 例如:
 *
 *
 * 输入:
 * A = [ 1, 2]
 * B = [-2,-1]
 * C = [-1, 2]
 * D = [ 0, 2]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 两个元组如下:
 * 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
 * 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 *
 *
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {

	return method1(A, B, C, D)
}

// 将四数之和转变为两数之和
func method1(A []int, B []int, C []int, D []int) int {
	hash, ans := make(map[int]int, 0), 0

	for _, a := range A{
		for _, b := range B{
			hash[a+b] ++
		}
	}

	for _, c := range C{
		for _, d := range D{
			if v, ok := hash[-(c+d)]; ok{
				ans += v
			}
		}
	}
	return ans
}
// @lc code=end
