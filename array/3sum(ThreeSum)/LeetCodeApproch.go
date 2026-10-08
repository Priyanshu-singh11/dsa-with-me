func threeSum(nums []int) [][]int {
    newArr := [][]int{}
    unique := make(map[[3]int]bool)
    sort.Ints(nums)
    for i:=0; i<len(nums);i++{
        left := i+1
        right := len(nums)-1
        for left<right{
            sum := nums[left]+nums[right]+nums[i]
            if sum<0{
                left++
            } else if sum>0{
                right--
            }else{
                triplet := [3]int{nums[i],nums[left],nums[right]}
                
                if !unique[triplet]{
                    unique[triplet] = true
                    
                newArr = append(newArr,triplet[:])
                }
                left++
                right--
            }
            
        }
    }
    return newArr
}






