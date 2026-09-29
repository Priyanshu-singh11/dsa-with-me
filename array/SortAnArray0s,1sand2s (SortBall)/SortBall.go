func sortColors(nums []int)  {
    var low,mid,high int = 0,0,len(nums)-1
    
    for mid<=high{
        if nums[mid]==0{
            nums[low],nums[mid]=nums[mid],nums[low]
            mid += 1
            low += 1
        }else if nums[mid]==1 {
            mid += 1
        }else{
            nums[high],nums[mid]=nums[mid],nums[high]
            high -= 1
        }
    
}
}
