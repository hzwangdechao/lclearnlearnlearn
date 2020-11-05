/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (49.94%)
 * Likes:    779
 * Dislikes: 0
 * Total Accepted:    84.8K
 * Total Submissions: 169.5K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 *
 *
 *
 * 进阶:
 *
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
 *
 *
 * 示例:
 *
 * LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得关键字 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得关键字 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 *
 *
 */

// @lc code=start

// 双向链表+hashmap
type LRUCache struct {
	Size int
	capacity int
	Cache map[int]*DLinkedNode
	Head, Tail *DLinkedNode
}

type DLinkedNode struct{
	Key, Value int
	Pre, Next *DLinkedNode
}

// 初始化双向链表
func initDLinkedNode(key, value int)*DLinkedNode  {
	return &DLinkedNode{
		Key: key,
		Value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Cache: map[int]*DLinkedNode{},
		Head: initDLinkedNode(0, 0),
		Tail: initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.Head.Next = l.Tail
	l.Tail.Pre = l.Head
	return l
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.Value
}


func (this *LRUCache) Put(key int, value int)  {
	_, ok := this.Cache[key]
	if !ok{
		// 在缓存中找不到的情况下， 先初始化一下新节点
		node := initDLinkedNode(key, value)
		// 将key>node放到缓存中
		this.Cache[key] = node
		// 将节点添加到链表的头部
		this.addToHead(node)
		this.Size ++

		// 如果cache的数量大于cache的容量
		if this.Size > this.capacity{
			// 删除链表和缓存中的尾部元素
			removed := this.removeTail()
			delete(this.Cache, removed.Key)
			this.Size --
		}

	}else{
		// 如果key在缓存中已经存在的情况下， 则将key移动到链表的头部, 并更新value
		node := this.Cache[key]
		node.Value = value
		this.moveToHead(node)
	}
}

// 链表相关的操作

// 删除双向链表中的节点
func (this *LRUCache)removeNode(node *DLinkedNode)  {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

// 将节点移动到双向链表的头部
func (this *LRUCache)addToHead(node *DLinkedNode)  {
	node.Pre = this.Head
	node.Next = this.Head.Next

	this.Head.Next.Pre = node
	this.Head.Next = node
}

// 将节点移动到链表的头部
func (this *LRUCache)moveToHead(node *DLinkedNode)  {
	// 先删除
	this.removeNode(node)
	// 再移动
	this.addToHead(node)
}

// 删除链表的尾部
func (this *LRUCache)removeTail()*DLinkedNode  {
	node := this.Tail.Pre
	this.removeNode(node)
	return node

}



/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
