package main

import "fmt"

//https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0004.Median-of-Two-Sorted-Arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// gia su nums1 la nho hon nums2
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // phia ben phai vach phan chia la Mid, ben trai la mid-1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // Đường phân chia trong nums1 bị vẽ quá nhiều và nó cần được di chuyển sang trái
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { //  Đường phân chia trong nums1 ít được vẽ hơn, nó cần được di chuyển sang bên phải
			low = nums1Mid + 1
		} else {
			// Tìm phép chia phù hợp, cần xuất ra kết quả cuối cùng
			// Chia thành 2 trường hợp số lẻ và số chẵn
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
