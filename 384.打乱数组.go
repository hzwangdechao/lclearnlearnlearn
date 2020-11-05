/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 *
 * https://leetcode-cn.com/problems/shuffle-an-array/description/
 *
 * algorithms
 * Medium (50.97%)
 * Likes:    75
 * Dislikes: 0
 * Total Accepted:    20.9K
 * Total Submissions: 40.1K
 * Testcase Example:  '["Solution","shuffle","reset","shuffle"]\n[[[1,2,3]],[],[],[]]'
 *
 * 打乱一个没有重复元素的数组。
 *
 *
 *
 * 示例:
 *
 * // 以数字集合 1, 2 和 3 初始化数组。
 * int[] nums = {1,2,3};
 * Solution solution = new Solution(nums);
 *
 * // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
 * solution.shuffle();
 *
 * // 重设数组到它的初始状态[1,2,3]。
 * solution.reset();
 *
 * // 随机返回数组[1,2,3]打乱后的结果。
 * solution.shuffle();
 *
 *
 */

// @lc code=start
type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {
	return Solution{nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums

}


/** Returns a random shuffling of the array. */

// 直接调用函数库打乱数组顺序
func (this *Solution) Shuffle1() []int {
	nums := make([]int, len(this.nums))
	copy(nums, this.nums)
	rand.Shuffle(len(nums), func (i, j int)  {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

// 创建新数据， 循环从数组中随机抽取元素添加到新数组中， 并且将已经抽取的元素删除
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.nums))
	buf := make([]int, len(this.nums))

	copy(buf, this.nums)
	for i, _:= range nums{
		j := rand.Intn(len(buf))
		nums[i] = buf[j]

		// 删除抽取掉的元素
		buf = append(buf[:j], buf[j+1:]...)

	}
	return nums

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
