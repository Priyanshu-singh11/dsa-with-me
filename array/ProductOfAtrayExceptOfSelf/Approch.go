func productExceptSelf(nums []int) []int {
    suffix := make([]int,len(nums))
    prefix := make([]int,len(nums))
    ans := make([]int,len(nums))
    suffix[len(nums)-1] = 1
    prefix[0] =1
    for i:=1;i<len(nums);i++{
        
            prefix [i] = prefix[i-1]*nums[i-1]
        
    }
    for j:=len(nums)-2;j>=0;j--{
        
            suffix[j] = suffix[j+1]*nums[j+1]
        
    }
    for m:=0;m<len(nums);m++{
        ans[m] = prefix[m]*suffix[m]
    }
    return ans
}
