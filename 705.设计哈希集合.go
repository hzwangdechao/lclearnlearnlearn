/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 *
 * https://leetcode-cn.com/problems/design-hashset/description/
 *
 * algorithms
 * Easy (56.84%)
 * Likes:    51
 * Dislikes: 0
 * Total Accepted:    16.2K
 * Total Submissions: 28.5K
 * Testcase Example:  '["MyHashSet","add","add","contains","contains","add","contains","remove","contains"]\n' +
  '[[],[1],[2],[1],[3],[2],[2],[2],[2]]'
 *
 * 不使用任何内建的哈希表库设计一个哈希集合
 *
 * 具体地说，你的设计应该包含以下的功能
 *
 *
 * add(value)：向哈希集合中插入一个值。
 * contains(value) ：返回哈希集合中是否存在这个值。
 * remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
 *
 *
 *
 * 示例:
 *
 * MyHashSet hashSet = new MyHashSet();
 * hashSet.add(1);        
 * hashSet.add(2);        
 * hashSet.contains(1);    // 返回 true
 * hashSet.contains(3);    // 返回 false (未找到)
 * hashSet.add(2);          
 * hashSet.contains(2);    // 返回 true
 * hashSet.remove(2);          
 * hashSet.contains(2);    // 返回  false (已经被删除)
 *
 *
 *
 * 注意：
 *
 *
 * 所有的值都在 [0, 1000000]的范围内。
 * 操作的总数目在[1, 10000]范围内。
 * 不要使用内建的哈希集合库。
 *
 *
 */

// @lc code=start

// Golang 单链表实现哈希集合
type Node struct{
	Val int
	Next *Node
}

type Bucket struct{
	Head *Node
}

func NewBucket()*Bucket  {
	return &Bucket{&Node{0, nil}}
}

func newNode(key int)*Node  {
	return &Node{
		Val: key,
		Next: nil,
	}
}

// 向单链表中添加节点
func (b *Bucket)insert(key int)  {
	cur := b.Head.Next
	node := newNode(key)

	// 如果在链表中找到了与key相同的节点， 则中断循环； 如果在链表中没有找到与key相同的节点， 则将新节点插入到链表的尾部
	for cur != nil{
		if cur.Val != key{
			if cur.Next != nil{
				cur = cur.Next
			}else{
				cur.Next = node
				break
			}
		}else{
			break
		}
	}
	if cur == nil{
		b.Head.Next = node
	}
}

// 在链表中删除节点
func (b *Bucket)delete(key int)  {
	cur := b.Head
	// 如果在链表中能找到相关的节点， 则删除指定的节点
	for cur.Next != nil{
		if cur.Next.Val == key{
			cur.Next = cur.Next.Next
			break
		}else{
			cur = cur.Next
		}
	}
}

// 检查链表中是否存在节点
func (b *Bucket)exists(key int)bool  {
	cur := b.Head.Next
	for cur != nil{
		if cur.Val == key{
			return true
		}else{
			cur = cur.Next
		}
	}
	return false
}

type MyHashSet struct {
	Set []*Bucket
}

const L = 769

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	// 新建L个链表的链表
	hs := MyHashSet{}
	for i:=0;i<L;i++{
		hs.Set = append(hs.Set, NewBucket())
	}
	fmt.Println(hs)
	return hs
}

func Hash(key int)int  {
	return key%L
}


func (this *MyHashSet) Add(key int)  {
	bucketIndex := Hash(key)
	this.Set[bucketIndex].insert(key)

}


func (this *MyHashSet) Remove(key int)  {
	bucketIndex := Hash(key)
	this.Set[bucketIndex].delete(key)
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bucketIndex := Hash(key)
	return this.Set[bucketIndex].exists(key)
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end
