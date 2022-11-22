package other

// https://leetcode.cn/problems/lfu-cache/
// 460. LFU 缓存
// 难度：困难
/*
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。

实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



示例：

输入：
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

解释：
// cnt(x) = 键 x 的使用计数
// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // 返回 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // 返回 -1（未找到）
lfu.get(3);      // 返回 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // 返回 4
                 // cache=[3,4], cnt(4)=2, cnt(3)=3
*/

type LFUListNode struct {
	Key  int
	Prev *LFUListNode
	Next *LFUListNode
}

type LFULinkedHashSet struct {
	hash map[int]*LFUListNode
	head *LFUListNode
	tail *LFUListNode
}

func NewLinkedHashSet() *LFULinkedHashSet {
	lhs := &LFULinkedHashSet{
		hash: make(map[int]*LFUListNode),
		head: &LFUListNode{},
		tail: &LFUListNode{},
	}

	lhs.head.Next = lhs.tail
	lhs.tail.Prev = lhs.head
	return lhs
}

func (l *LFULinkedHashSet) Remove(key int) bool {
	node, ok := l.hash[key]
	if !ok {
		return false
	}

	l.removeNode(node)
	delete(l.hash, key)
	return true
}

func (l *LFULinkedHashSet) RemoveTail() (int, bool) {
	if l.IsEmpty() {
		return 0, false
	}

	node := l.tail.Prev
	return node.Key, l.Remove(node.Key)
}

func (l *LFULinkedHashSet) Add(key int) bool {
	if _, ok := l.hash[key]; ok {
		// already exists
		return false
	}

	node := &LFUListNode{key, nil, nil}
	l.hash[key] = node
	l.pushFront(node)
	return true
}

func (l *LFULinkedHashSet) IsEmpty() bool {
	return l.head.Next == l.tail
}

func (l *LFULinkedHashSet) pushFront(node *LFUListNode) {
	node.Next = l.head.Next
	node.Prev = l.head
	l.head.Next.Prev = node
	l.head.Next = node
}

func (l *LFULinkedHashSet) removeNode(node *LFUListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Prev, node.Next = nil, nil
}

type LFUCache struct {
	capacity   int                       // capacity
	keyToVal   map[int]int               // key -> value
	keyToFreq  map[int]int               // key -> frequency
	freqToKeys map[int]*LFULinkedHashSet // frequency -> hset
	minFreq    int
}

func Constructor(capacity int) *LFUCache {
	lfu := &LFUCache{
		capacity:   capacity,
		keyToVal:   make(map[int]int),
		keyToFreq:  make(map[int]int),
		freqToKeys: make(map[int]*LFULinkedHashSet),
		minFreq:    0,
	}
	return lfu
}

func (this *LFUCache) Get(key int) int {
	v, ok := this.keyToVal[key]
	if !ok {
		// key不存在
		return -1
	}

	// 更新key的访问频率
	this.increaseFreq(key)
	return v
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}

	// key已存在，修改对应的val
	if _, ok := this.keyToVal[key]; ok {
		this.keyToVal[key] = value
		// 更新key的频率
		this.increaseFreq(key)
		return
	}

	// key不存在，插入
	if len(this.keyToVal) >= this.capacity {
		this.removeMinFreqKey()
	}

	// 插入KV表
	this.keyToVal[key] = value

	// 插入KF表
	this.keyToFreq[key] = 1

	// 插入FK表
	this.addFreqToKeys(1, key)

	// 插入新的key，最小的freq肯定是1
	this.minFreq = 1
}

// 移除使用频率最低的key
func (this *LFUCache) removeMinFreqKey() {
	keyList, ok := this.freqToKeys[this.minFreq]
	if !ok {
		return
	}

	// 从最小频率列表中移除最先插入的元素
	deleteKey, _ := keyList.RemoveTail()
	if keyList.IsEmpty() {
		delete(this.freqToKeys, this.minFreq)
		// removeMinFreqKey只在put的时候调用，put时会更新minFreq，所以这里不需要更新
	}

	// KV表中移除key
	delete(this.keyToVal, deleteKey)

	// KF表中移除key
	delete(this.keyToFreq, deleteKey)
}

func (this *LFUCache) increaseFreq(key int) {
	freq := this.keyToFreq[key]

	// 更新KF表
	newFreq := freq + 1
	this.keyToFreq[key] = newFreq

	// 更新FK表
	// 将key从freq对应的列表中删除
	this.freqToKeys[freq].Remove(key)

	// 将key加入到newFreq对应的列表中
	this.addFreqToKeys(newFreq, key)

	// freq对应的列表为空，移除这个freq
	if this.freqToKeys[freq].IsEmpty() {
		delete(this.freqToKeys, freq)
		if freq == this.minFreq {
			this.minFreq = newFreq
		}
	}
}

func (this *LFUCache) addFreqToKeys(freq int, key int) {
	if _, ok := this.freqToKeys[freq]; !ok {
		this.freqToKeys[freq] = NewLinkedHashSet()
	}
	this.freqToKeys[freq].Add(key)
}
