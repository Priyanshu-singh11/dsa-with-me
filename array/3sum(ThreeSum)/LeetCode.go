func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    newArr := [][]int{}
    for i:=0;i<len(nums);i++{
        if i>0 && nums[i]==nums[i-1]{
            continue 
        }
        j:=i+1
        k := len(nums)-1
        for j<k{
            sum := nums[i]+nums[j]+nums[k]
            if sum>0{
                k--
            }else if sum<0{
                j++
            }else{
                arr := [3]int {nums[i],nums[j],nums[k]}
                newArr = append(newArr,arr[:])
                j++
                k--
                for j<k && nums[j]==nums[j-1]{
                j++
            }
            for j<k && nums[k]==nums[k+1]{
                k--
            }
            }
            
        }
    }
    return newArr
}






