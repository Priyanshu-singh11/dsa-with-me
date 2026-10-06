
func threeSum(nums []int) [][]int {
    newArr := [][]int{}
    unique := make(map[[3]int]bool)
    for i:=0;i<len(nums);i++{
        a := nums[i]
        seen := make(map[int]bool)
        for j:=i+1;j<len(nums);j++{
            b := nums[j]
            c := -(a+b)
            if seen[c]{
                triplet := [3]int{a,b,c}
                sort.Ints(triplet[:])
                if !unique[triplet]{
                    unique[triplet] = true
                    newArr = append(newArr,triplet[:])
                }
                
            }
            seen[b] = true 
        }
    }
    return newArr
}
