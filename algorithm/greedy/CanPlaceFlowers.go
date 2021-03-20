package greedy

// https://leetcode-cn.com/problems/can-place-flowers/
// 题目：605. 种花问题
// 难度：简单
func canPlaceFlowers(flowerbed []int, n int) bool {
	var plant int
	for i := 0; i < len(flowerbed); i++ {
		if (flowerbed[i] == 0) &&
			(i-1 < 0 || flowerbed[i-1] == 0) &&
			(i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			plant++
		}
	}
	return plant >= n
}
