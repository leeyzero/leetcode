## 二分搜索

二分搜索也常被称为二分法或折半查找，每次查找时通过将待查找的区间分成两部分并取一部分继续查找，将查找的复杂度大减少。对于一个长度为O(N)的数组，二分搜索的时间复杂度为O(logN).

二分搜索虽然算法思路比较简单，但如果细节没有处理好，往往得不到正确的答案。其中最关键的细节在于你采用的区间表示方式，主要有两种：
- 左闭右闭，即[left, right]
- 左闭右开，即[left, right）

怎么理解呢？对于一个长度为n的数据，采用左闭右闭区间表示为：[0, n-1]；采用左闭右开区间表示为[0, n)。

## 算法框架
```
func binarySearch(nums []int, target int) int {
    left, right := 0, ...
    for ... {
        mid := left + (right - left)/2
        if nums[mid] == target {
            ...
        } else if nums[mid]) < target {
            left = ...
        } else if nums[mid] > target {
            right = ...
        }
    }
    return ...
}
```

其中 `...` 部分即为一些细节处理，而这些细节处理，跟你采用的区间表示有关系。

## 基本二分搜索
**左闭右闭区间表示**
```
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums) - 1 // 注意左闭右闭区间表示方式
    for left <= right { // 左闭右闭的结束条件是left == right + 1
         mid := left + (right - left)/2
        if nums[mid] == target {
            return mid // 找到target
        } else if nums[mid]) < target {
            left = mid + 1  // 收缩左边界
        } else if nums[mid] > target {
            right = mid - 1 // 收缩右边界
        }
    }
    // 未找到
    return -1
}
```

**左闭右开区间表示**
```
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums) // 注意左闭右开区间表示方式
    for left < right { // 左闭右开的结束条件是left == right
        mid := left + (right - left)/2
        if nums[mid] == target {
            return mid // 找到target
        } else if nums[mid]) < target {
            left = mid + 1  // 收缩左边界
        } else if nums[mid] > target {
            right = mid // 注意：由于采用左闭右开的表示，收缩右边界，right取值为mid
        }
    }

    // 注意：从结束条件left == right来看，如果left还在合法区间内，我们将会漏掉left元素的处理，例如区间[2,2]未非空区间，需要打个补丁：
    if left < len(nums) && nums[left] == target {
        return left
    }

    // 未找到
    return -1
}
```

## 左边界二分搜索
左边界二分搜索是二分搜索的一种变体，用于搜索有序数组中重复元素的右边界。例如，给定一个有序数组 `nums = [1, 3, 3, 3, 4], target = 3`，返回 `target` 的左边界，即索引1。

如果理解了上面基本的二分搜索，确定区间的表示方式后，惟一的区别在于**当找到target元素的索引后，需要收缩右边界，这样才能让左右指针不断收敛到target的最左边界。**

**左闭右闭区间表示**
```
func binarySearchLeftBound(nums []int, target int) int {
    left, right := 0, len(nums) - 1 // 注意左闭右闭区间表示方式
    for left <= right { // 结束条件是 left == right + 1
        mid := left + (right - left)/2
        if nums[mid] == target {
            right = mid - 1 // 注意：找到target后，需要收缩右边界。由于区间采用左闭右闭表示，right为mid-1
        } else if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        }
    }

    // 从结束条件来看，如果left是一个合法索引，并且索引的值为target，则left即为结果。
    if left < len(nums) && nums[left] == target {
        return left
    }

    // 未找到
    return -1
}
```

**左闭右开区间表示**
```
func binarySearchLeftBound(nums []int, target int) int {
    left, right := 0, len(nums) // 注意左闭右开区间表示方式
    for left < right { // 结束条件是 left == right
        mid := left + (right - left)/2
        if nums[mid] == target {
            right = mid // 注意：找到target后，需要收缩右边界。由于区间采用左闭右开表示，right为mid
        } else if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid // 注意：由于区间采用左闭右开表示，right为mid
        }
    }

    // 从结束条件来看，如果找到target，则left必定为最左侧的target索引。
    if left < len(nums) && nums[left] == target {
        return left
    }

    // 未找到
    return -1
}
```

## 右边界二分搜索
跟左边界二分搜索类似，用于搜索有序数组中重复元素的右边界。例如，给定一个有序数组 `nums = [1, 3, 3, 3, 4], target = 3`，返回 `target` 的右边界，即索引3。

如果理解了上面基本的二分搜索，确定区间的表示方式后，惟一的区别在于**当找到target元素的索引后，需要收缩左边界，这样才能让左右指针不断收敛到target的最右边界。**

**左闭右闭区间表示**
```
func binarySearchRightBound(nums []int, target int) int {
    left, right := 0, len(nums) - 1 // 注意左闭右闭区间表示方式
    for left <= right { // 结束条件是 left == right + 1
        mid := left + (right - left)/2
        if nums[mid] == target {
            left = mid + 1 // 找到target后，需要收缩左边界。
        } else if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        }
    }

    // 从结束条件来看，如果left大于0，那么left-1索引的值为target，则left-1索引的值即为target的右边界。
    if left > 0 && nums[left-1] == target {
        return left-1
    }

    // 未找到
    return -1
}
```

为什么最后的结果是left-1呢？可以这样理解：当`nums[mid] == target`时，left赋值为`mid+1`，如果这个时候退出循环，那么`nums[left]`的值一定不等于`target`，而`left-1`则可能是等于`target`的。

**左闭右开区间表示**
```
func binarySearchRightBound(nums []int, target int) int {
    left, right := 0, len(nums) // 注意左闭右开区间表示方式
    for left < right { // 结束条件是 left == right
        mid := left + (right - left)/2
        if nums[mid] == target {
            left = mid + 1 // 找到target后，需要收缩左边界。
        } else if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid // 注意：由于采用左闭右开区间表示，right收缩时为mid（不包括mid）
        }
    }

    // 从结束条件来看，如果left大于0，那么left-1索引的值为target，则left-1索引的值即为target的右边界。
    if left > 0 && nums[left-1] == target {
        return left-1
    }

    // 未找到
    return -1
}
```

## 常见题目
 - [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)
 - [35. 搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/)
 - [剑指 Offer 53 - I. 在排序数组中查找数字I](https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/)
 - [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)
 - [81. 搜索旋转排序数组 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)

## 参考资料
- [labuladong的算法小抄](https://item.jd.com/12759911.html)

- [Binary Search](https://www.geeksforgeeks.org/binary-search/)

- [Binary search algorithm](https://en.wikipedia.org/wiki/Binary_search_algorithm)

