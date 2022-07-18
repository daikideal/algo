package main

func maxSubArray(nums []int) int {
	currentSum, maxSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum+nums[i], nums[i])
		maxSum = max(currentSum, maxSum)

		// fmt.Printf("when i:%d, currentSum:%d maxSum:%d\n", i, currentSum, maxSum)
	}
	return maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
