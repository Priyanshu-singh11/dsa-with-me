
func MooreVotingAlgorithm(nums []int) int {
	count := 0
	ans := 0
	//   1, 3, 3, 2, 3, 2
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			ans = nums[i]
		}
		if ans == nums[i] {
			count++
		} else {
			count--
		}
	}

	//if value not exist then it verify
	count2 := 0
	for _, num := range nums {
		if num == ans {
			count2++
		}
	}
	if count2 > len(nums)/2 {
		return ans
	}
	return -1
}
