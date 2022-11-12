## 双指针

双指针是一项简单但却非常有用的技术，特别是在排序数组中搜索目标对象时特别有效。
指针在计算机中是指它引用的对象。双指针，顾名思义，就是用两个指针跟进数组或字符串中的两个不同的对象，以解决目标问题。

双指针技术有两种比较常见的模式：
```
1、两个指针分别指向头和尾，相向移动，直到它们相遇。
2、两个指针均从头同向移动，一个移动的慢，另一个指针移动的快。
```

## 经验问题
[有序数组中和为s的两个数字](https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/) 
难度：简单 
解题思路：首尾指针求和比较，相向移动，直到找到目标值或指针相遇。 

[判断链表是否有环](https://leetcode-cn.com/problems/linked-list-cycle/) 
难度：简单 
解题思路：快慢指针，相向移动，直到它们相遇或链表结尾。 

[删除有序数据中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/) 
难度：简单 
解题思路：双指针，分别跟踪无重复项和遍历数组，lo, hi 初始化为-1, 0。hi遍历数组，出现hi和lo不同时，lo向前移动，并将hi的值赋值给lo，
直到遍历结束。

[合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/) 
难度：简单 
解题思路：如果允许开辟空间，会更简单一些，用两个指针分别指向两个数组的起始位置，比较将较小的值取出放入新空间中。但题目中，第一个数组有
足够的空间，要求将结果放入第一个数组中。解题需要有点变通，用两个指针指向末尾，向前移动，并将较大的值放至第一个数组的末尾，直至遍历完两数组。

[调整数组顺序使奇数位于偶数前面](https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/) 
难度：简单 
解题思路：这个题跟[删除有序数据中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)类似，用两
个指针分别跟踪已处理的下标和遍历数组，直到遍历完成。

[数组中重复的数字](https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/) 
难度：简单 
解题思路：这个题目比较直观的解法是使用hash，但题目有一个隐含的条件是数组中的所有数字都在0-n~1范围内。所以可以利用这个特性，让数组本身作为一个
hash表，表的索引（下标）和值是相同的特性来解决。这样可以让空间复杂度做到O(1)。我们一个指针从数组头开始，跟踪目前已经处理的索引（下标），如果索引
和指向的值相等，说明该值在正确的位置，移动指针。如果该值不在正确的位置且该值作为索引指向的值已经在正确的位置，说明重复，否则将该值放到正确的位置。

[环形链表II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

难度：中等 
解题思路：这种题如果不了解其原理，在短短的20分钟内，其实很难做出来。基本思路是使用快慢指针和一些技巧。

解法一：先使用快慢指针判断链表是否有环，如果有环，那否一定能找到相遇点。当快慢指针第一次相遇时，如果慢指针走了k步，那么快指针走了2k步。
快指针比慢指针多走的k步，一定是从慢指针处绕环的整数倍。假设环入口点到相遇点的距离为m，则链表头到环入口点的距离则为k-m。看起来问题转换为让
指针从链表头开始走k-m步就找到了环的入口点。但k-m我们并不知道是多少？巧的是，快指针此时再走k-m步，也到达了环的入口处（快慢指针相遇时，快
指针多走了k步，k-m正好环的入口点）。所以我们可以将慢指针指向链表头，然后跟快指针一起移动，当他们相遇时，相遇点就是环的入口点。

解法二：仍然是使用快慢指针先找到相遇点（如果存在）。判断相遇点后，我们可以计算出环的长度（再从相遇点跑一圈）k，假设链表头到环入口点的距离
为m（不包括入口点），则链表长度m+k。我们先让快指针走k步，然后再让慢指针从头开始走跟快指针一起走，当快指针到达入口点时，走了m+k步，此时慢
指针刚好走了m步，也就是也到达了环的入口点，它们刚好相遇。所以快慢指针的相遇就是环的入口点。


## 参考资料
[1] [The Two Pointer Technique](https://algodaily.com/lessons/using-the-two-pointer-technique/python)

[2] [Two-pointer technique](https://leetcode.com/articles/two-pointer-technique/)

[3] [Two Pointers Technique](https://www.geeksforgeeks.org/two-pointers-technique/)
