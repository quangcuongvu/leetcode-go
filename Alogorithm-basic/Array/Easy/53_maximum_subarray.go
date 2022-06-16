package main

//https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0053.Maximum-Subarray/
//https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}
func main() {
	nums := []int{2,-3,5,-3}
	result := maxSubArray(nums)
	println(result)

}
