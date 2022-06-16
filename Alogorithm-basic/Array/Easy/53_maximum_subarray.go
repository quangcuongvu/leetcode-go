package main

//https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0053.Maximum-Subarray/
//https://leetcode.com/problems/maximum-subarray/
//video: https://www.youtube.com/watch?v=t1bxDeSz2Rg

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func maxSubArray2(nums []int) int {
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

func maxSubArrayYoutube(nums []int) int {
	result := nums[0]
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp = max(tmp+nums[i], nums[i])
		result = max(result, tmp)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, -3, 5, -3}
	result := maxSubArrayYoutube(nums)
	println(result)

}
