package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	l := removeElement(nums, val)
	fmt.Println(l, nums)
}

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
func moveZeroes(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], append(nums[i+1:], nums[i])...)
			i--
			l--
		}
	}
}

func removeElement(nums []int, val int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			l--
		}
	}
	return len(nums)
}
