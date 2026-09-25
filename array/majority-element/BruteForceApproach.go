
func BruteForceApproach(nums []int) int {
	// you can keep count := 0 then you want to do j:= 0
	for i := 0; i < len(nums); i++ {
		count := 1 //inside because it reset counts
		//{3, 3, 3, 1, 1, 1, 3}
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}
