package algorithm

// LRU: least recent used
// 条件要求Get和Put满足O(1)时间复杂度，看起来哈希可以满足，但在移除最久未使用的数据或调整最近访问的数据时
// 仅仅靠hash做不到，但链表可以做到在O(1)时间增加或删除数据，将hash和链表结合起来可以满足需求。
// 对于链表节点，有两个问题需要回答：
// 1. 节点中为什要存储key
// 2. 为什要使用双向链表
// 对于第一个问题是为了解决方向查找，当在淘汰最久未使用的节点时，如果每次往链表尾插入的话，最久未使用的节点就
// 是头节点，除了删除头节点外，还用删除hash中的对应key的节点，如果链表节点中没有存储key的话，是找不到hash中
// 对应的节点的。
// 对于第二个问题，当在访问一个已存在的节点时，需要将其调整到尾部，双向链表可以方便删除链表中的任意节点。
type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

// hash用于存储key到LinkNode的映射
// head和tail是哨兵节点，便于处理边界问题
// capacity是缓存大小
type LRUCache struct {
	hash     map[int]*LinkNode
	head     *LinkNode
	tail     *LinkNode
	capacity int
}

// 构造函数除了创建节点外，还要将head和tail连接起来
func NewLRUCache(capacity int) *LRUCache {
	obj := LRUCache{
		hash:     map[int]*LinkNode{},
		head:     &LinkNode{},
		tail:     &LinkNode{},
		capacity: capacity,
	}
	obj.head.next = obj.tail
	obj.tail.pre = obj.head
	return &obj
}

// 先检查key是否存在，不存在返回-1
// key存在，需要调整key对应的节点为最近使用
func (this *LRUCache) Get(key int) int {
	if node, ok := this.get(key); ok {
		this.makeRecent(node)
		return node.val
	}
	return -1
}

// 先检查key是否存在，key存在的情况下，除了更新key对应的值外，还需要将key对应的节点调整为最近使用节点
// key不存在，还需要检出缓存是否已满，缓存满时需要淘汰最久未使用的节点。
// 最后将节点加入到队尾，并缓存到hash表中
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.get(key); ok {
		node.val = value
		this.makeRecent(node)
		return
	}

	if this.isFull() {
		this.removeLeastRecentUsed()
	}
	this.put(&LinkNode{key, value, nil, nil})
}

func (this *LRUCache) get(key int) (*LinkNode, bool) {
	if node, ok := this.hash[key]; ok {
		return node, true
	}
	return nil, false
}

func (this *LRUCache) put(node *LinkNode) {
	if node == nil {
		return
	}

	this.pushTail(node)
	this.hash[node.key] = node
}

func (this *LRUCache) removeLeastRecentUsed() {
	if front, ok := this.popFront(); ok {
		delete(this.hash, front.key)
	}
}

// 将节点node重新添加到队尾
func (this *LRUCache) makeRecent(node *LinkNode) {
	this.removeNode(node)
	this.pushTail(node)
}

func (this *LRUCache) removeNode(node *LinkNode) {
	if node == nil {
		return
	}

	// pre -> next
	node.pre.next = node.next
	// next -> pre
	node.next.pre = node.pre
	// nil <- node -> nil
	node.pre, node.next = nil, nil
}

// 将节点添加到队尾
func (this *LRUCache) pushTail(node *LinkNode) {
	if node == nil {
		return
	}

	// last <- node -> tail
	node.pre, node.next = this.tail.pre, this.tail
	// last -> node
	this.tail.pre.next = node
	// node <- tail
	this.tail.pre = node
}

// 弹出队头节点
func (this *LRUCache) popFront() (*LinkNode, bool) {
	if this.isEmpty() {
		return nil, false
	}

	front := this.head.next
	this.removeNode(front)
	return front, true
}

func (this *LRUCache) isEmpty() bool {
	return len(this.hash) == 0
}

func (this *LRUCache) isFull() bool {
	return len(this.hash) >= this.capacity
}
