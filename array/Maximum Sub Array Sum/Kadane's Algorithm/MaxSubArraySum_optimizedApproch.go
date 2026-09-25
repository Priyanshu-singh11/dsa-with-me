
func maxSubArray(nums []int) int {
    //var maxElem = math.MinInt
    maxElem := int(math.Inf(-1))
    currElem := 0
    for _,num := range nums{
        currElem += num
        if currElem>maxElem{
            maxElem = currElem
        }
        if currElem<0{
            currElem = 0
        }
    }
    return maxElem
}
