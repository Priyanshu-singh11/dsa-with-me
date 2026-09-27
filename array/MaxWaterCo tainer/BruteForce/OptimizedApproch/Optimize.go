func maxArea(height []int) int {
    var left,right int = 0,len(height)-1
    var maxWater int = 0
    for left<right {
        var h,w int = 0,right-left
        if(height[left]>height[right]){
            h = height[right]
            right -= 1
        }else{
            h = height[left]
            left += 1
        }
        var area int = w*h
        if(area>maxWater){
            maxWater = area
        }
    }
    return maxWater
}
