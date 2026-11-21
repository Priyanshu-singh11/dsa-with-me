func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    newArr := [][]int{}
    for i:=0;i<len(nums);i++{
        if i>0 && nums[i]==nums[i-1]{
            continue 
        }
        for j:=i+1;j<len(nums);{
            p := j+1
            q := len(nums)-1
            for p<q{
                sum := nums[i]+nums[j]+nums[p]+nums[q]
                if sum>target{
                    q--
                }else if sum<target{
                    p++
                }else{
                    arr := []int{nums[i],nums[j],nums[p],nums[q]}
                    newArr = append(newArr,arr[:])
                    p++
                    q--
                    for p<q && nums[p]==nums[p-1]{
                        p++
                    }
                    //q-- loop not required 
                    for p<q && nums[q]==nums[q+1]{
                        q--
                    }
                }
            }
            j++
            for j<len(nums) && nums[j]==nums[j-1]{
                j++
            }
        }
    }
    return newArr
}
