/*
 * @lc app=leetcode.cn id=974 lang=golang
 *
 * [974] 和可被 K 整除的子数组
 *
 * https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/description/
 *
 * algorithms
 * Medium (37.80%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    16.7K
 * Total Submissions: 37.8K
 * Testcase Example:  '[4,5,0,-2,-3,1]\n5'
 *
 * 给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
 *
 *
 *
 * 示例：
 *
 * 输入：A = [4,5,0,-2,-3,1], K = 5
 * 输出：7
 * 解释：
 * 有 7 个子数组满足其元素之和可被 K = 5 整除：
 * [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2,
 * -3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 30000
 * -10000 <= A[i] <= 10000
 * 2 <= K <= 10000
 *
 *
 */

// @lc code=start
func subarraysDivByK(A []int, K int) int {
	// hash表， 逐一统计
	// m := map[int]int{0:1}  // 记录余数的次数
	// sum , res := 0, 0
	// for _, x := range A{
	// 	sum += x
	// 	modulus := (sum % K + K) % K  // +K 是为了防止出现负数
	// 	fmt.Println(modulus)
	// 	res += m[modulus]
	// 	fmt.Println(res)
	// 	fmt.Println(m[modulus])


	// 	m[modulus]++
	// 	fmt.Println("---")
	// }
	// return res

	// hash表， 单次统计
	m := map[int]int{0:1}
	sum, res := 0, 0
	for _, x := range A{
		sum += x
		modulus := (sum %K + K) %K

		m[modulus] ++

	}
	fmt.Println(m)
	for _, x := range m{
		fmt.Println(x)
		res += x * (x-1)/2
		fmt.Println(res)
	}

	return res



}
// @lc code=end
