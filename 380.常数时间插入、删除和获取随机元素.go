/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 *
 * https://leetcode-cn.com/problems/insert-delete-getrandom-o1/description/
 *
 * algorithms
 * Medium (48.59%)
 * Likes:    161
 * Dislikes: 0
 * Total Accepted:    14.8K
 * Total Submissions: 30.4K
 * Testcase Example:  '["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]\n' +
  '[[],[1],[2],[2],[],[1],[2],[]]'
 *
 * 设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
 *
 *
 * insert(val)：当元素 val 不存在时，向集合中插入该项。
 * remove(val)：元素 val 存在时，从集合中移除该项。
 * getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
 *
 *
 * 示例 :
 *
 *
 * // 初始化一个空的集合。
 * RandomizedSet randomSet = new RandomizedSet();
 *
 * // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
 * randomSet.insert(1);
 *
 * // 返回 false ，表示集合中不存在 2 。
 * randomSet.remove(2);
 *
 * // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
 * randomSet.insert(2);
 *
 * // getRandom 应随机返回 1 或 2 。
 * randomSet.getRandom();
 *
 * // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
 * randomSet.remove(1);
 *
 * // 2 已在集合中，所以返回 false 。
 * randomSet.insert(2);
 *
 * // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
 * randomSet.getRandom();
 *
 *
 */

// @lc code=start
type RandomizedSet struct {
	rMap map[int]int
	rSlice []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	// 如果存在的话返回false
	if _, ok := this.rMap[val]; ok{
		return false
	}

	// 不存在的话则返回map和slice
	this.rSlice = append(this.rSlice, val)
	this.rMap[val] = len(this.rSlice) - 1
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	// 不存在的话返回false
	if _, ok := this.rMap[val]; !ok{
		return false
	}

	// 把index与队尾元素调换， 并且更新map中的index对应关系
	index := this.rMap[val]
	this.rSlice[index] = this.rSlice[len(this.rSlice)-1]
	this.rMap[this.rSlice[index]] = index

	// 删除队尾元素， 并且删除map中的信息
	this.rSlice = this.rSlice[:len(this.rSlice)-1]
	delete(this.rMap, val)
	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.rSlice) == 0{
		return -1
	}

	index := rand.Intn(len(this.rSlice))
	return this.rSlice[index]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
