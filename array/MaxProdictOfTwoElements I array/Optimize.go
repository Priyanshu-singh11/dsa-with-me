
func maxProduct(nums []int) int {
    var largest,secondLargest = 0,0
    for i:=0;i<len(nums);i++{
        if nums[i]>=largest{
            secondLargest = largest
            largest = nums[i]
        }else if nums[i]>secondLargest{
            secondLargest = nums[i]
        }
    }
    return (largest-1)*(secondLargest-1)
}
