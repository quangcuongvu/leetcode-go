package main

//https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0035.Search-Insert-Position/
//https://leetcode.com/problems/search-insert-position/

import "fmt"

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] >= target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	result := searchInsert(nums, target)
	fmt.Println(result)
}
