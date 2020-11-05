/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 *
 * https://leetcode-cn.com/problems/design-skiplist/description/
 *
 * algorithms
 * Hard (58.91%)
 * Likes:    32
 * Dislikes: 0
 * Total Accepted:    1.9K
 * Total Submissions: 3.3K
 * Testcase Example:  '["Skiplist","add","add","add","search","add","search","erase","erase","search"]\r\n' +
  '[[],[1],[2],[3],[0],[4],[1],[0],[1],[1]]\r'
 *
 * 不使用任何库函数，设计一个跳表。
 *
 * 跳表是在 O(log(n))
 * 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，并且跳表的代码长度相较下更短，其设计思想与链表相似。
 *
 * 例如，一个跳表包含 [30, 40, 50, 60, 70, 90]，然后增加 80、45 到跳表中，以下图的方式操作：
 *
 *
 * Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons
 *
 * 跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。跳表的每一个操作的平均时间复杂度是
 * O(log(n))，空间复杂度是 O(n)。
 *
 * 在本题中，你的设计应该要包含这些函数：
 *
 *
 * bool search(int target) : 返回target是否存在于跳表中。
 * void add(int num): 插入一个元素到跳表。
 * bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num
 * ，删除其中任意一个即可。
 *
 *
 * 了解更多 : https://en.wikipedia.org/wiki/Skip_list
 *
 * 注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。
 *
 * 样例:
 *
 * Skiplist skiplist = new Skiplist();
 *
 * skiplist.add(1);
 * skiplist.add(2);
 * skiplist.add(3);
 * skiplist.search(0);   // 返回 false
 * skiplist.add(4);
 * skiplist.search(1);   // 返回 true
 * skiplist.erase(0);    // 返回 false，0 不在跳表中
 * skiplist.erase(1);    // 返回 true
 * skiplist.search(1);   // 返回 false，1 已被擦除
 *
 *
 * 约束条件:
 *
 *
 * 0 <= num, target <= 20000
 * 最多调用 50000 次 search, add, 以及 erase操作。
 *
 *
 */

// @lc code=start
const MaxLevel = 7

type Node struct{
	Value int  // 存储值
	Prev *Node  // 同层的上一个节点
	Next *Node  // 同层的下一个节点
	Down *Node  // 下层同节点
}

type Skiplist struct {
	Level int  // 层级
	HeadNodeArr []*Node // 存储每一层的头节点
}


func Constructor() Skiplist {
	arr := make([]*Node, MaxLevel)
	return Skiplist{0, arr}
}

// original
func (this *Skiplist) Search(target int) bool {
	node := this.hasNode(target)
	if node == nil{
		return false
	}
	return true
}

func (this *Skiplist) hasNode(target int) *Node  {
	fmt.Println(this.HeadNodeArr)
	if this.Level >= 0{
		// 只有在层级大于等于0的时候才进行循环判断， 如果层级小于0说明没有任何数据
		level := this.Level
		node := this.HeadNodeArr[level].Next


		for node != nil{
			if node.Value == target{
				// 如果节点的值等于目标值， 返回node
				return node
			}else if node.Value > target{
				// 如果当前节点的值大于目标值， 则回退到前一个节点并进入下一层
				if node.Prev.Down == nil{
					if level - 1 >= 0{
						// 从头开始
						node = this.HeadNodeArr[level-1].Next
					}else{
						node = nil
					}
				}else{
					node = node.Prev.Down
				}
				level -= 1
			}else if node.Value < target{
				// 如果当前节点的值小于目标值， 就进入下一个节点， 如果下一个节点为空的话， 说明本层已经查询完毕， 从下一层的头部开始
				node = node.Next
				if node == nil{
					level -= 1
					if level >= 0{
						// 如果不是最底层就进入到下一层
						node = this.HeadNodeArr[level].Next
					}
				}
			}
		}
	}
	return nil
}

// original
func (this *Skiplist) Add(num int)  {
	// 如果已经存在相同的元素， 则不进行添加的操作
	if this.Search(num){
		return
	}
	headNodeInsertPositionArr := make([]*Node, MaxLevel)
	if this.Level >= 0{
		level := this.Level
		node := this.HeadNodeArr[level].Next

		for node != nil && level >= 0{
			if node.Value > num {
				// 如果当前的值大于传入的值， 则返回上一个节点并进入下一层
				headNodeInsertPositionArr[level] = node.Prev
				if node.Prev.Down == nil{
					if level - 1 >= 0{
						node = this.HeadNodeArr[level-1].Next
					}else{
						node = nil
					}
				}else{
					node = node.Prev.Down
				}
				level -= 1
			}
		}else if node.Value < num{
			if node.Next == nil{
				headNodeInsertPositionArr[level] = node
				level -= 1
				if level >= 0{
					node = this.HeadNodeArr[level].Next
				}
			}else{
				node = node.Next
			}
		}
	}
    this.InsertValue(value, headNodeInsertPositionArr)
}



// original
func (this *Skiplist) Erase(num int) bool {
	// 先查找是否存在目标值
	node := this.hasNode(num)
	if node == nil{
		return false
	}

	for node != nil{
		prevNode := node.Prev
		nextNode := node.Next

		prevNode.Next = nextNode // 将上一个节点的next指向下一个节点
		if nextNode != nil{
			nextNode.Prev = prevNode  // 将下一个节点的prev指向上一个节点
		}
		// 向下继续进行操作
		node = node.Down
	}
	return true
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end
