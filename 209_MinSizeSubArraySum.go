package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLen := n + 1
	lhs := 0
	sum := 0
	for rhs := 0; rhs < n; rhs++ {
		sum += nums[rhs]
		for sum >= target {
			if rhs-lhs+1 < minLen {
				minLen = rhs - lhs + 1
			}
			sum -= nums[lhs]
			lhs++
		}
	}
	return minLen % (n + 1)
}

/*
	minSubArrayLen(6, []int{0, 2, 1, 5})

	target: 6

	nums: 	0		2		1		5
									i,		j

	currSum: 5

	minLen:  1


*/
