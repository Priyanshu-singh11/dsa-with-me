
func maxSubArray(nums []int) int {
    var maxElem int = math.MinInt
    
    for i:=0;i<len(nums);i++{
        currElem := 0
        for j:=i;j<len(nums);j++{
            currElem += nums[j]
            if currElem>maxElem{
                maxElem = currElem
            }
        }
    }
    return maxElem
}
