package main

func sum(nums ...int) int {
	output := 0
	for i := range nums {
		output += nums[i]
	}
	return output
}
